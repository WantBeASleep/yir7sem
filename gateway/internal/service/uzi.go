package service

import (
	"context"
	"yir/gateway/internal/custom"
	"yir/gateway/internal/entity/uzimodel"
	"yir/gateway/internal/entity/uzimodel/uzidto"
	"yir/gateway/internal/pb/uzipb"
	"yir/gateway/internal/service/mapper/uzimap"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UziService struct {
	Client uzipb.UziAPIClient
}

// НЕ ИСПОЛЬЗУЕТСЯ
// func (u *UziService) InsertUzi(ctx context.Context, in *uzidto.Uzi, image *entity.Image) error {
// 	// перегоняем структурки
// 	req := &uzipb.Report{
// 		Uzi: &uzipb.Uzi{
// 			Url:        in.UziInfo.URL,
// 			Projection: in.UziInfo.Projection,
// 			PatientId:  in.UziInfo.PatientID.String(),
// 			DeviceId:   int64(in.UziInfo.DeviceID),
// 		},
// 	}
// 	req.Formations = make([]*uzipb.Formation, 0, len(in.Formations))
// 	req.Images = make([]*uzipb.Image, 0, len(in.Images))
// 	req.Segments = make([]*uzipb.Segment, 0, len(in.Segments))
// 	for _, v := range in.Formations {
// 		req.Formations = append(req.Formations,
// 			&uzipb.Formation{
// 				Id: v.Id.String(),
// 				Tirads: &uzipb.Tirads{
// 					Tirads_1: v.Tirads.Tirads1,
// 					Tirads_2: v.Tirads.Tirads2,
// 					Tirads_3: v.Tirads.Tirads3,
// 					Tirads_4: v.Tirads.Tirads4,
// 					Tirads_5: v.Tirads.Tirads5,
// 				},
// 				Ai: v.Ai,
// 			},
// 		)
// 	}
// 	for _, v := range in.Images {
// 		req.Images = append(req.Images,
// 			&uzipb.Image{
// 				Id:   v.ID.String(),
// 				Url:  v.URL,
// 				Page: int64(v.Page),
// 			},
// 		)
// 	}
// 	for _, v := range in.Segments {
// 		req.Segments = append(req.Segments,
// 			&uzipb.Segment{
// 				Id:          v.Id.String(),
// 				FormationId: v.FormationID.String(),
// 				ImageId:     v.ImageID.String(),
// 				ContorUrl:   v.ContorURL,
// 				Tirads: &uzipb.Tirads{
// 					Tirads_1: v.Tirads.Tirads1,
// 					Tirads_2: v.Tirads.Tirads2,
// 					Tirads_3: v.Tirads.Tirads3,
// 					Tirads_4: v.Tirads.Tirads4,
// 					Tirads_5: v.Tirads.Tirads5,
// 				},
// 			},
// 		)
// 	}
// 	// отдаем в узи
// 	if _, err := u.Client.InsertUzi(ctx, req); err != nil {
// 		custom.Logger.Error(
// 			"failed to insert uzi",
// 			zap.Error(err),
// 		)
// 		return err
// 	}
// 	return nil
// }

func (u *UziService) CreateUzi(ctx context.Context, in *uzimodel.Uzi) (string, error) {
	req := &uzipb.CreateUziRequest{
		Uzi: &uzipb.UziRequest{
			Url:        in.URL,
			Projection: in.Projection,
			PatientId:  in.PatientID.String(),
			DeviceId:   int64(in.DeviceID),
		},
	}
	data, err := u.Client.CreateUzi(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed to create uzi info",
			zap.Error(err),
		)
		return "", err
	}

	return data.Id, nil
}

func (u *UziService) UpdateUzi(ctx context.Context, in *uzimodel.Uzi) error {
	req := &uzipb.UpdateUziRequest{
		Uzi: &uzipb.UziRequest{
			Url:        in.URL,
			Projection: in.Projection,
			PatientId:  in.PatientID.String(),
			DeviceId:   int64(in.DeviceID),
		},
	}
	if _, err := u.Client.UpdateUzi(ctx, req); err != nil {
		custom.Logger.Error(
			"failed to update uzi info",
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (u *UziService) CreateFormationWithSegments(ctx context.Context, uziID string, formations *uzidto.FormationWithSegments) error {
	// ебать тут нахуй дерьма
	req := &uzipb.CreateFormationWithSegmentsRequest{
		UziId: uziID,
		Formation: &uzipb.FormationWithNestedSegmentsRequest{
			Segments: make([]*uzipb.SegmentNestedRequest, 0, len(formations.Segments)),
			Tirads:   uzimap.TiradsToPBTirads(formations.Formation.Tirads),
			Ai:       formations.Formation.Ai,
		},
	}

	for _, segment := range formations.Segments {
		req.Formation.Segments = append(
			req.Formation.Segments,
			&uzipb.SegmentNestedRequest{
				ImageId:   segment.ImageID.String(),
				ContorUrl: segment.ContorURL,
				Tirads:    uzimap.TiradsToPBTirads(segment.Tirads),
			},
		)
	}
	if _, err := u.Client.CreateFormationWithSegments(ctx, req); err != nil {
		custom.Logger.Error(
			"failed to insert formation with segments",
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (u *UziService) UpdateFormation(ctx context.Context, formation *uzidto.Formation) error {
	req := &uzipb.UpdateFormationRequest{
		Formation: &uzipb.FormationRequest{
			Ai: formation.Ai,
			Tirads: &uzipb.Tirads{
				Tirads_23: formation.Tirads.Tirads23,
				Tirads_4:  formation.Tirads.Tirads4,
				Tirads_5:  formation.Tirads.Tirads5,
			},
		},
	}
	if _, err := u.Client.UpdateFormation(ctx, req); err != nil {
		custom.Logger.Error(
			"failed to update formation",
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (u *UziService) GetUziByID(ctx context.Context, uziID string) (*uzimodel.Uzi, error) {
	req := &uzipb.Id{
		Id: uziID,
	}
	data, err := u.Client.GetUzi(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed to get uzi by id",
			zap.Error(err),
		)
		return nil, err
	}
	resp := &uzimodel.Uzi{}
	// перенос uzi info
	resp.ID, err = uuid.Parse(data.Id)
	if err != nil {
		custom.Logger.Error(
			"failed parsing uzi uuid",
			zap.Error(err),
		)
		return nil, err
	}
	resp.PatientID, err = uuid.Parse(data.PatientId)
	if err != nil {
		custom.Logger.Error(
			"failed parsing uzi patient uuid",
			zap.Error(err),
		)
		return nil, err
	}
	resp.DeviceID = int(data.DeviceId)
	resp.Projection = data.Projection
	resp.URL = data.Url
	return resp, nil
}

func (u *UziService) GetImageWithFormationsSegments(ctx context.Context, imageID string) (*uzidto.ImageWithSegmentsFormations, error) {
	req := &uzipb.Id{
		Id: imageID,
	}
	data, err := u.Client.GetImageWithFormationsSegments(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed to get image segments and formations info",
			zap.Error(err),
		)
		return nil, err
	}
	id, err := uuid.Parse(data.Image.Id)
	if err != nil {
		custom.Logger.Error(
			"failed parsing image uuid",
			zap.Error(err),
		)
		return nil, err
	}
	uziID, err := uuid.Parse(data.Image.Id) /// НЕТУ ПОЛЯ UziId в data, а структура Image имеет это поле
	if err != nil {
		custom.Logger.Error(
			"failed parsing uzi uuid",
			zap.Error(err),
		)
		return nil, err
	}
	resp := &uzidto.ImageWithSegmentsFormations{
		Image: &uzimodel.Image{
			ID:    id,
			URL:   data.Image.Url,
			Page:  int(data.Image.Page),
			UziID: uziID,
		},
	}
	resp.Formations = make([]uzidto.Formation, 0, len(data.Formations))
	for _, v := range data.Formations {
		id, err := uuid.Parse(v.Id)
		if err != nil {
			custom.Logger.Error(
				"failed parsing formation uuid",
				zap.Error(err),
			)
			return nil, err
		}
		resp.Formations = append(
			resp.Formations,
			uzidto.Formation{
				Id: id,
				Ai: v.Ai,
				Tirads: &uzimodel.Tirads{
					Tirads23: v.Tirads.Tirads_23,
					Tirads4:  v.Tirads.Tirads_4,
					Tirads5:  v.Tirads.Tirads_5,
				},
			},
		)
	}
	// segments
	resp.Segments = make([]uzidto.Segment, 0, len(data.Segments))
	for _, v := range data.Segments {
		id, err := uuid.Parse(v.Id)
		if err != nil {
			custom.Logger.Error(
				"failed parsing segment id uuid",
				zap.Error(err),
			)
			return nil, err
		}
		FormationId, err := uuid.Parse(v.FormationId)
		if err != nil {
			custom.Logger.Error(
				"failed parsing segment formation uuid",
				zap.Error(err),
			)
			return nil, err
		}
		ImageId, err := uuid.Parse(v.ImageId)
		if err != nil {
			custom.Logger.Error(
				"failed parsing segment image uuid",
				zap.Error(err),
			)
			return nil, err
		}
		resp.Segments = append(
			resp.Segments,
			uzidto.Segment{
				Id:          id,
				ImageID:     ImageId,
				FormationID: FormationId,
				ContorURL:   v.ContorUrl,
				Tirads:      uzimap.PBTiradsToTirads(v.Tirads),
			},
		)
	}
	return resp, nil
}

func (u *UziService) GetFormationWithSegments(ctx context.Context, formationID string) (*uzidto.FormationWithSegments, error) {
	req := &uzipb.Id{
		Id: formationID,
	}
	data, err := u.Client.GetFormationWithSegments(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed to get formation with segments",
			zap.Error(err),
		)
		return nil, err
	}
	FormationID, err := uuid.Parse(formationID)
	if err != nil {
		custom.Logger.Error(
			"failed to parse formation uuid",
			zap.Error(err),
		)
		return nil, err
	}

	resp := &uzidto.FormationWithSegments{
		Formation: &uzidto.Formation{
			Id:     FormationID,
			Ai:     data.Formation.Ai,
			Tirads: uzimap.PBTiradsToTirads(data.Formation.Tirads),
		},
	}
	resp.Segments = make([]uzidto.Segment, 0, len(data.Segments))
	for _, v := range data.Segments {
		id, err := uuid.Parse(v.Id)
		if err != nil {
			custom.Logger.Error(
				"failed parsing segment id uuid",
				zap.Error(err),
			)
			return nil, err
		}
		FormationId, err := uuid.Parse(v.FormationId)
		if err != nil {
			custom.Logger.Error(
				"failed parsing segment formation uuid",
				zap.Error(err),
			)
			return nil, err
		}
		ImageId, err := uuid.Parse(v.ImageId)
		if err != nil {
			custom.Logger.Error(
				"failed parsing segment image uuid",
				zap.Error(err),
			)
			return nil, err
		}
		resp.Segments = append(
			resp.Segments,
			uzidto.Segment{
				Id:          id,
				ImageID:     ImageId,
				FormationID: FormationId,
				ContorURL:   v.ContorUrl,
				Tirads:      uzimap.PBTiradsToTirads(v.Tirads),
			},
		)
	}
	return resp, nil
}
func (u *UziService) GetDeviceList(ctx context.Context) ([]uzimodel.Device, error) {
	data, err := u.Client.GetDeviceList(ctx, &emptypb.Empty{})
	if err != nil {
		custom.Logger.Error(
			"failed to get device list",
			zap.Error(err),
		)
		return nil, err
	}
	resp := make([]uzimodel.Device, 0, len(data.Devices))
	for _, v := range data.Devices {
		resp = append(
			resp,
			uzimodel.Device{
				ID:   int(v.Id),
				Name: v.Name,
			},
		)
	}
	return resp, nil
}
