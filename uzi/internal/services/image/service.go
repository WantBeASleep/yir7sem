package image

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	"uzi/internal/adapters"
	uzisplittedpb "uzi/internal/generated/broker/produce/uzisplitted"

	"uzi/internal/domain"
	"uzi/internal/repository"
	"uzi/internal/repository/entity"
	"uzi/internal/services/splitter"

	"github.com/google/uuid"
)

type Service interface {
	GetUziImages(ctx context.Context, uziID uuid.UUID) ([]domain.Image, error)
	GetImageSegmentsWithNodes(ctx context.Context, id uuid.UUID) ([]domain.Node, []domain.Segment, error)
	SplitUzi(ctx context.Context, uziID uuid.UUID) error
}

type service struct {
	dao     repository.DAO
	adapter adapters.Adapter
}

func New(
	dao repository.DAO,
	adapter adapters.Adapter,
) Service {
	return &service{
		dao:     dao,
		adapter: adapter,
	}
}

func (s *service) CreateImages(ctx context.Context, images []domain.Image) ([]uuid.UUID, error) {
	ids := make([]uuid.UUID, 0, len(images))
	for i := range images {
		images[i].Id = uuid.New()
		ids = append(ids, images[i].Id)
	}

	imagesDB := make([]entity.Image, 0, len(images))
	for _, v := range images {
		imagesDB = append(imagesDB, entity.Image{}.FromDomain(v))
	}

	if err := s.dao.NewImageQuery(ctx).InsertImages(imagesDB); err != nil {
		return nil, fmt.Errorf("insert images: %w", err)
	}
	return ids, nil
}

func (s *service) GetUziImages(ctx context.Context, uziID uuid.UUID) ([]domain.Image, error) {
	images, err := s.dao.NewImageQuery(ctx).GetImagesByUziID(uziID)
	if err != nil {
		return nil, fmt.Errorf("get images by uzi_id: %w", err)
	}

	return entity.Image{}.SliceToDomain(images), nil
}

func (s *service) GetImageSegmentsWithNodes(ctx context.Context, id uuid.UUID) ([]domain.Node, []domain.Segment, error) {
	segments, err := s.dao.NewSegmentQuery(ctx).GetSegmentsByImageID(id)
	if err != nil {
		return nil, nil, fmt.Errorf("get segments by image_id: %w", err)
	}

	// TODO: переделать на запросе без JOIN
	nodes, err := s.dao.NewNodeQuery(ctx).GetNodesByImageID(id)
	if err != nil {
		return nil, nil, fmt.Errorf("get nodes by image_id: %w", err)
	}

	return entity.Node{}.SliceToDomain(nodes), entity.Segment{}.SliceToDomain(segments), nil
}

// TODO: возвращать отсюда ID
// выгрузить из s3
// засплитить
// загрузить в psql
// загрузить в s3
// написать в kafka
func (s *service) SplitUzi(ctx context.Context, uziID uuid.UUID) error {
	fileRepo := s.dao.NewFileRepo()

	exists, err := s.dao.NewUziQuery(ctx).CheckExist(uziID)
	if err != nil {
		return fmt.Errorf("check exists uzi: %w", err)
	}
	if !exists {
		return errors.New("uzi doesnt exist")
	}

	file, closer, err := fileRepo.GetFileViaTemp(ctx, filepath.Join(uziID.String(), uziID.String()))
	if err != nil {
		return fmt.Errorf("get file via temp: %w", err)
	}
	defer closer()

	splitterSrv := splitter.New()
	splitted, err := splitterSrv.SplitToPng(file)
	if err != nil {
		return fmt.Errorf("split img to png: %w", err)
	}

	images := make([]domain.Image, len(splitted))
	for i := range images {
		images[i].UziID = uziID
		images[i].Page = i + 1
	}

	// TODO: сделать транзакцию
	ids, err := s.CreateImages(ctx, images)
	if err != nil {
		return fmt.Errorf("create Images: %w", err)
	}

	for i, v := range ids {
		if err := fileRepo.LoadFile(ctx, filepath.Join(uziID.String(), v.String(), v.String()), splitted[i]); err != nil {
			return fmt.Errorf("load file to S3: %w", err)
		}
	}

	if err := s.adapter.BrokerAdapter.SendUziSplitted(&uzisplittedpb.UziSplitted{
		UziId:   uziID.String(),
		PagesId: uuid.UUIDs(ids).Strings(),
	}); err != nil {
		return fmt.Errorf("send to uzisplitted topic: %w", err)
	}

	return nil
}
