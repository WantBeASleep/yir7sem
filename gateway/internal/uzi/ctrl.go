package uzi

import (
	"context"
	pb "yir/gateway/rpc/uzi"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Ctrl struct {
	pb.UnimplementedUziAPIServer

	client pb.UziAPIClient
}

func (c *Ctrl) CreateUzi(ctx context.Context, req *pb.CreateUziRequest) (*pb.Id, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

// GetUzi godoc
//	@Summary		Получить УЗИ по ID
//	@Description	Возвращает информацию о УЗИ по указанному ID.
//	@Tags			Uzi
//	@Produce		json
//	@Param			uzi_id	path		string	true	"ID УЗИ"
//	@Success		200		{object}	Uzi		"Данные УЗИ"
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/uzi/{uzi_id} [get]
func (c *Ctrl) GetUzi(ctx context.Context, req *pb.Id) (*pb.UziReponse, error) {
	return c.client.GetUzi(ctx, req)
}

// GetReport godoc
//	@Summary		Получить Репорт по ID
//	@Description	Возвращает информацию о Репорте по указанному ID.
//	@Tags			Uzi
//	@Produce		json
//	@Param			uzi_id	path		string	true	"ID УЗИ"
//	@Success		200		{object}	Report	"Данные УЗИ"
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/report/{uzi_id} [get]
func (c *Ctrl) GetReport(ctx context.Context, req *pb.Id) (*pb.Report, error) {
	return c.client.GetReport(ctx, req)
}

// UpdateUzi godoc
//	@Summary		Обновить УЗИ
//	@Description	Обновляет существующую запись УЗИ по переданным данным.
//	@Tags			Uzi
//	@Accept			json
//	@Produce		json
//	@Param			uzi_id	path		string	true	"ID УЗИ"
//	@Param			body	body		UziReq	true	"Метаданные УЗИ"
//	@Success		200		{object}	Uzi		"Успешное выполнение"
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/uzi/{uzi_id} [patch]
func (c *Ctrl) UpdateUzi(ctx context.Context, req *pb.UpdateUziRequest) (*pb.UziReponse, error) {
	return c.client.UpdateUzi(ctx, req)
}

// GetUziImageWithSegments godoc
//	@Summary		Получить изображение УЗИ с сегментами
//	@Description	Возвращает изображение УЗИ и информацию о сегментах по указанному ID изображения.
//	@Tags			Uzi
//	@Produce		json
//	@Param			image_id	path		string						true	"ID изображения"
//	@Success		200			{object}	ImageWithSegmentsFormations	"Изображение с сегментами"
//	@Failure		500			{string}	string						"Internal error"
//	@Router			/uzi/image/segments/{image_id} [get]
func (c *Ctrl) GetImageWithFormationsSegments(ctx context.Context, req *pb.Id) (*pb.ImageWithFormationsSegments, error) {
	return c.client.GetImageWithFormationsSegments(ctx, req)
}

// TODO: узел и не тока здесь
// CreateFormationWithSegments godoc
//	@Summary		Добавить формацию с сегментами
//	@Description	Создает формацию с сегментами, привязанную к указанному uziID.
//	@Tags			Uzi
//	@Accept			json
//	@Produce		json
//	@Param			uzi_id	path		string						true	"ID УЗИ"
//	@Param			body	body		FormationNestedSegmentReq	true	"Формация с сегментами"
//	@Success		200		{object}	FormationWithSegmentsIDs	"ID элементов"
//	@Failure		500		{string}	string						"Internal error"
//	@Router			/formation/segments/{uzi_id} [post]
func (c *Ctrl) CreateFormationWithSegments(ctx context.Context, req *pb.CreateFormationWithSegmentsRequest) (*pb.CreateFormationWithSegmentsResponse, error) {
	return c.client.CreateFormationWithSegments(ctx, req)
}

// GetFormationWithSegments godoc
//	@Summary		Получить формацию с сегментами
//	@Description	Возвращает формацию с сегментами по указанному ID формации.
//	@Tags			Uzi
//	@Produce		json
//	@Param			formation_id	path		string					true	"ID формации"
//	@Success		200				{object}	FormationWithSegments	"Данные формации с сегментами"
//	@Failure		500				{string}	string					"Internal error"
//	@Router			/formation/segments/{formation_id} [get]
func (c *Ctrl) GetFormationWithSegments(ctx context.Context, req *pb.Id) (*pb.FormationWithSegments, error) {
	return c.client.GetFormationWithSegments(ctx, req)
}

// TODO: это поменять, убрать grpc-gateway для узи
// UpdateFormation godoc
//	@Summary		Обновить формацию
//	@Description	Обновляет существующую формацию по переданному ID.
//	@Tags			Uzi
//	@Accept			json
//	@Produce		json
//	@Param			formation_id	path		string			true	"ID формации"
//	@Param			body			body		FormationReq	true	"Данные формации"
//	@Success		200				{object}	Formation		"Успешное выполнение"
//	@Failure		500				{string}	string			"Internal error"
//	@Router			/uzi/formation/{formation_id} [patch]
func (c *Ctrl) UpdateFormation(ctx context.Context, req *pb.UpdateFormationRequest) (*pb.FormationResponse, error) {
	return c.client.UpdateFormation(ctx, req)
}

// UpdateSegment godoc
//	@Summary		Обновить сегмент
//	@Description	Обновляет существующую сегмент по переданному ID.
//	@Tags			Uzi
//	@Accept			json
//	@Produce		json
//	@Param			segment_id	path		string			true	"ID формации"
//	@Param			body		body		SegmentUpdate	true	"Данные формации"
//	@Success		200			{object}	Segment			"Успешное выполнение"
//	@Failure		500			{string}	string			"Internal error"
//	@Router			/uzi/segment/{segment_id} [patch]
func (c *Ctrl) UpdateSegment(ctx context.Context, req *pb.UpdateSegmentRequest) (*pb.SegmentResponse, error) {
	return c.client.UpdateSegment(ctx, req)
}

// GetDeviceList godoc
//	@Summary		Получить список устройств
//	@Description	Возвращает список доступных устройств для УЗИ.
//	@Tags			Uzi
//	@Produce		json
//	@Success		200	{array}		Device	"Список устройств"
//	@Failure		500	{string}	string	"Internal error"
//	@Router			/uzi/device/list [get]
func (c *Ctrl) GetDeviceList(ctx context.Context, req *empty.Empty) (*pb.GetDeviceListResponse, error) {
	return c.client.GetDeviceList(ctx, req)
}
