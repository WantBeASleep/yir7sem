package uzi

import (
	"context"
	"log"
	pb "yir/gateway/rpc/uzi"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Ctrl struct {
	pb.UnimplementedUziAPIServer

	client pb.UziAPIClient
}

func NewCtrl(
	client pb.UziAPIClient,
) *Ctrl {
	return &Ctrl{
		client: client,
	}
}

// CreateUzi godoc
//
//	@Summary		ЗагрузитьUzi
//	@Description	Да пиздец.
//	@Tags			Uzi
//	@Produce		json
//	@Failure		500	{string}	string	"Internal error"
//	@Router			/uzi/create [post]
func (c *Ctrl) CreateUzi(ctx context.Context, req *pb.UziRequest) (*pb.Id, error) {
	log.Println("Called CreateUzi")
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

// GetUzi godoc
//
//	@Summary		Получить УЗИ по ID
//	@Description	Возвращает информацию о УЗИ по указанному ID.
//	@Tags			Uzi
//	@Produce		json
//	@Param			uzi_id	path		string				true	"ID УЗИ"
//	@Success		200		{object}	grpcapi.UziReponse	"Данные УЗИ"
//	@Failure		500		{string}	string				"Internal error"
//	@Router			/uzi/{uzi_id} [get]
func (c *Ctrl) GetUzi(ctx context.Context, req *pb.Id) (*pb.UziReponse, error) {
	log.Println("Called GetUzi")
	return c.client.GetUzi(ctx, req)
}

// GetReport godoc
//
//	@Summary		Получить Репорт по ID
//	@Description	Возвращает информацию о Репорте по указанному ID.
//	@Tags			Uzi
//	@Produce		json
//	@Param			uzi_id	path		string			true	"ID УЗИ"
//	@Success		200		{object}	grpcapi.Report	"Данные УЗИ"
//	@Failure		500		{string}	string			"Internal error"
//	@Router			/report/{uzi_id} [get]
func (c *Ctrl) GetReport(ctx context.Context, req *pb.Id) (*pb.Report, error) {
	log.Println("Called GetReport")
	return c.client.GetReport(ctx, req)
}

// UpdateUzi godoc
//
//	@Summary		Обновить УЗИ
//	@Description	Обновляет существующую запись УЗИ по переданным данным.
//	@Tags			Uzi
//	@Accept			json
//	@Produce		json
//	@Param			uzi_id	path		string						true	"ID УЗИ"
//	@Param			body	body		grpcapi.UziUpdateRequest	true	"Метаданные УЗИ"
//	@Success		200		{object}	grpcapi.UziReponse			"Успешное выполнение"
//	@Failure		500		{string}	string						"Internal error"
//	@Router			/uzi/{uzi_id} [put]
func (c *Ctrl) UpdateUzi(ctx context.Context, req *pb.UpdateUziRequest) (*pb.UziReponse, error) {
	log.Println("Called UpdateUzi")
	return c.client.UpdateUzi(ctx, req)
}

// GetUziImageWithSegments godoc
//
//	@Summary		Получить изображение УЗИ с сегментами
//	@Description	Возвращает изображение УЗИ и информацию о сегментах по указанному ID изображения.
//	@Tags			Uzi
//	@Produce		json
//	@Param			image_id	path		string								true	"ID изображения"
//	@Success		200			{object}	grpcapi.ImageWithFormationsSegments	"Изображение с сегментами"
//	@Failure		500			{string}	string								"Internal error"
//	@Router			/uzi/image/withsegments/{image_id} [get]
func (c *Ctrl) GetImageWithFormationsSegments(ctx context.Context, req *pb.Id) (*pb.ImageWithFormationsSegments, error) {
	log.Println("Called GetImageWithFormationsSegments")
	return c.client.GetImageWithFormationsSegments(ctx, req)
}

// CreateFormationWithSegments godoc
//
//	@Summary		Добавить формацию с сегментами
//	@Description	Создает формацию с сегментами, привязанную к указанному uziID.
//	@Tags			Uzi
//	@Accept			json
//	@Produce		json
//	@Param			uzi_id	path		string										true	"ID УЗИ"
//	@Param			body	body		grpcapi.FormationWithNestedSegmentsRequest	true	"Формация с сегментами"
//	@Success		200		{object}	grpcapi.CreateFormationWithSegmentsResponse	"ID элементов"
//	@Failure		500		{string}	string										"Internal error"
//	@Router			/uzi/formation/withsegments/{uzi_id} [post]
func (c *Ctrl) CreateFormationWithSegments(ctx context.Context, req *pb.CreateFormationWithSegmentsRequest) (*pb.CreateFormationWithSegmentsResponse, error) {
	log.Println("Called CreateFormationWithSegments")
	return c.client.CreateFormationWithSegments(ctx, req)
}

// GetFormationWithSegments godoc
//
//	@Summary		Получить формацию с сегментами
//	@Description	Возвращает формацию с сегментами по указанному ID формации.
//	@Tags			Uzi
//	@Produce		json
//	@Param			formation_id	path		string							true	"ID формации"
//	@Success		200				{object}	grpcapi.FormationWithSegments	"Данные формации с сегментами"
//	@Failure		500				{string}	string							"Internal error"
//	@Router			/uzi/formation/withsegments/{formation_id} [get]
func (c *Ctrl) GetFormationWithSegments(ctx context.Context, req *pb.Id) (*pb.FormationWithSegments, error) {
	log.Println("Called GetFormationWithSegments")
	return c.client.GetFormationWithSegments(ctx, req)
}

// TODO: это поменять, убрать grpc-gateway для узи
// UpdateFormation godoc
//
//	@Summary		Обновить формацию
//	@Description	Обновляет существующую формацию по переданному ID.
//	@Tags			Uzi
//	@Accept			json
//	@Produce		json
//	@Param			formation_id	path		string						true	"ID формации"
//	@Param			body			body		grpcapi.FormationRequest	true	"Данные формации"
//	@Success		200				{object}	grpcapi.FormationResponse	"Успешное выполнение"
//	@Failure		500				{string}	string						"Internal error"
//	@Router			/uzi/formation/{formation_id} [put]
func (c *Ctrl) UpdateFormation(ctx context.Context, req *pb.UpdateFormationRequest) (*pb.FormationResponse, error) {
	log.Println("Called UpdateFormation")
	return c.client.UpdateFormation(ctx, req)
}

// UpdateSegment godoc
//
//	@Summary		Обновить сегмент
//	@Description	Обновляет существующую сегмент по переданному ID. (Мы его не тестили)
//	@Tags			Uzi
//	@Accept			json
//	@Produce		json
//	@Param			segment_id	path		string					true	"ID формации"
//	@Param			body		body		grpcapi.SegmentRequest	true	"Данные формации"
//	@Success		200			{object}	grpcapi.SegmentResponse	"Успешное выполнение"
//	@Failure		500			{string}	string					"Internal error"
//	@Router			/uzi/segment/{segment_id} [put]
func (c *Ctrl) UpdateSegment(ctx context.Context, req *pb.UpdateSegmentRequest) (*pb.SegmentResponse, error) {
	log.Println("Called UpdateSegment")
	return c.client.UpdateSegment(ctx, req)
}

// GetDeviceList godoc
//
//	@Summary		Получить список устройств
//	@Description	Возвращает список доступных устройств для УЗИ.
//	@Tags			Uzi
//	@Produce		json
//	@Success		200	{array}		grpcapi.Device	"Список устройств"
//	@Failure		500	{string}	string			"Internal error"
//	@Router			/uzi/device/list [get]
func (c *Ctrl) GetDeviceList(ctx context.Context, req *empty.Empty) (*pb.GetDeviceListResponse, error) {
	log.Println("Called GetDeviceList")
	return c.client.GetDeviceList(ctx, req)
}
