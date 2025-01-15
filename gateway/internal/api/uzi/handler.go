package uzi

// TODO: большая проблема: то что рисуем на выход в сваггер != тому что туда реально уходит (уходит GRPC)
// TODO: сделать выход контуров через json nyy
import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	adapters "gateway/internal/adapters"
	"gateway/internal/domain"

	uziuploadpb "gateway/internal/generated/broker/produce/uziupload"
	uzipb "gateway/internal/generated/grpc/client/uzi"
	"gateway/internal/repository"

	"github.com/gorilla/mux"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	adapter adapters.Adapter
	dao     repository.DAO
}

func New(
	adapter adapters.Adapter,
	dao repository.DAO,
) *Handler {
	return &Handler{
		adapter: adapter,
		dao:     dao,
	}
}

// PostUzi загружает узи на обработку
//
//	@Summary		Загружает узи на обработку
//	@Description	Загружает узи на обработку
//	@Tags			uzi
//	@Produce		json
//	@Param			token		header		string	true	"access_token"
//	@Param			file		formData	file	true	"uzi file. (обязательно с .tiff/.png)"
//	@Param			projection	formData	string	true	"проекция узи"
//	@Param			patient_id	formData	string	true	"id пациента"
//	@Param			device_id	formData	string	true	"id узи апапапапарата"
//	@Success		200			{string}	string	"molodec"
//	@Failure		500			{string}	string	"Internal Server Error"
//	@Router			/uzi/uzis [post]
func (h *Handler) PostUzi(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projection := r.FormValue("projection")
	patientID := r.FormValue("patient_id")
	deviceID, _ := strconv.Atoi(r.FormValue("device_id"))

	uziResp, err := h.adapter.UziAdapter.CreateUzi(ctx, &uzipb.CreateUziIn{
		Projection: projection,
		PatientId:  patientID,
		DeviceId:   int64(deviceID),
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	file, meta, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
	defer file.Close()
	ext := filepath.Ext(meta.Filename)

	// TODO: заюзать библу
	mime, err := domain.ParseFormatFromExt(ext)
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	err = h.dao.NewFileRepo().LoadFile(ctx, filepath.Join(uziResp.Id, uziResp.Id), domain.File{Format: mime, Buf: file})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	// TODO: нужна тотальная сага тут
	if err := h.adapter.BrokerAdapter.SendUziUpload(&uziuploadpb.UziUpload{UziId: uziResp.Id}); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(uziResp.Id); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// TODO: проверить крайние случае, если что то не приходит например(неправильный id)
// TODO: убрать echographic из ответа на обновление
// PatchUzi Обновляет узи
//
//	@Summary		Обновляет узи
//	@Description	Обновляет узи
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string		true	"access_token"
//	@Param			id		path		string		true	"uzi_id"
//	@Param			body	body		PatchUziIn	true	"обновляемые значения"
//	@Success		200		{object}	PatchUziOut	"uzi"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/uzi/uzis/{id} [patch]
func (h *Handler) PatchUzi(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var req PatchUziIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	resp, err := h.adapter.UziAdapter.UpdateUzi(ctx, &uzipb.UpdateUziIn{
		Id:         id,
		Projection: req.Projection,
		Checked:    req.Checked,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp.Uzi); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// PatchEchographics Обновляет эхографику
//
//	@Summary		Обновляет эхографику
//	@Description	Обновляет эхографику
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string					true	"access_token"
//	@Param			id		path		string					true	"uzi_id"
//	@Param			body	body		PatchEchographicsIn		true	"обновляемые значения"
//	@Success		200		{object}	PatchEchographicsOut	"echographic"
//	@Failure		500		{string}	string					"Internal Server Error"
//	@Router			/uzi/echographics/{id} [patch]
func (h *Handler) PatchEchographics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var req PatchEchographicsIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	resp, err := h.adapter.UziAdapter.UpdateEchographic(ctx, &uzipb.UpdateEchographicIn{
		Echographic: &uzipb.Echographic{
			Id:              id,
			LeftLobeLength:  req.LeftLobeLength,
			LeftLobeWidth:   req.LeftLobeWidth,
			LeftLobeThick:   req.LeftLobeThick,
			LeftLobeVolum:   req.LeftLobeVolum,
			RightLobeLength: req.RightLobeLength,
			RightLobeWidth:  req.RightLobeWidth,
			RightLobeThick:  req.RightLobeThick,
			RightLobeVolum:  req.RightLobeVolum,
			GlandVolum:      req.GlandVolum,
			Isthmus:         req.Isthmus,
			Struct:          req.Struct,
			Echogenicity:    req.Echogenicity,
			RegionalLymph:   req.RegionalLymph,
			Vascularization: req.Vascularization,
			Location:        req.Location,
			Additional:      req.Additional,
			Conclusion:      req.Conclusion,
		},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp.Echographic); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetUzi получает uzi
//
//	@Summary		получает uiz
//	@Description	получает uiz
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string		true	"access_token"
//	@Param			id		path		string		true	"uzi_id"
//	@Success		200		{object}	GetUziOut	"uzi"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/uzi/uzis/{id} [get]
func (h *Handler) GetUzi(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.UziAdapter.GetUzi(ctx, &uzipb.GetUziIn{Id: id})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
	// TODO: понять почему тут узи возвращается без эхографикой, а тут с
	// TODO: подумать над content-tpye в ответе(посмотреть в каком порядке выставлять функции для ответа)
	if err := json.NewEncoder(w).Encode(resp.Uzi); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetPatientUzi Получить узи пациента
//
//	@Summary		Получить узи пациента
//	@Description	Получить узи пациента
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Param			id		path		string	true	"patient_id"
//	@Success		200		{object}	GetPatientUziOut
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/uzi/patient/{id}/uzis [get]
func (h *Handler) GetPatientUzi(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	res, err := h.adapter.UziAdapter.GetPatientUzis(ctx, &uzipb.GetPatientUzisIn{
		PatientId: id,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	w.WriteHeader(200)
}

// GetEchographics получает uzi
//
//	@Summary		получает эхографику uzi
//	@Description	получает эхографику uzi
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string				true	"access_token"
//	@Param			id		path		string				true	"uzi_id"
//	@Success		200		{object}	GetEchographicsOut	"echographics"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/uzi/echographics/{id} [get]
func (h *Handler) GetEchographics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.UziAdapter.GetEchographic(ctx, &uzipb.GetEchographicIn{Id: id})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
	// TODO: понять почему тут узи возвращается без эхографикой, а тут с
	// TODO: подумать над content-tpye в ответе(посмотреть в каком порядке выставлять функции для ответа)
	if err := json.NewEncoder(w).Encode(resp.Echographic); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetUziImages получает id картинок uzi
//
//	@Summary		получает списк id кадров uzi
//	@Description	получает списк id кадров uzi
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string			true	"access_token"
//	@Param			id		path		string			true	"uzi_id"
//	@Success		200		{object}	GetUziImagesOut	"images"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/uzi/uzis/{id}/images [get]
func (h *Handler) GetUziImages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.UziAdapter.GetUziImages(ctx, &uzipb.GetUziImagesIn{UziId: id})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp.Images); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetUziNodeSegments получит ноды и сегменты на указанном изображении
//
//	@Summary		получит ноды и сегменты на указанном изображении
//	@Description	получит ноды и сегменты на указанном изображении
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string					true	"access_token"
//	@Param			id		path		string					true	"image_id"
//	@Success		200		{object}	GetUziNodeSegmentsOut	"nodes&&segments"
//	@Failure		500		{string}	string					"Internal Server Error"
//	@Router			/uzi/images/{id}/nodes-segments [get]
func (h *Handler) GetUziNodeSegments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.UziAdapter.GetImageSegmentsWithNodes(
		ctx,
		&uzipb.GetImageSegmentsWithNodesIn{Id: id},
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetUziDevice получит список uzi апппапапратов
//
//	@Summary		получит список uzi апппапапратов
//	@Description	получит список uzi апппапапратов
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string			true	"access_token"
//	@Success		200		{object}	GetUziDeviceOut	"uzi аппараты"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/uzi/devices [get]
func (h *Handler) GetUziDevices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp, err := h.adapter.UziAdapter.GetDeviceList(ctx, &emptypb.Empty{})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// PostNodes добавить узел с сегментами
//
//	@Summary		добавить узел с сегментами
//	@Description	добавить узел с сегментами
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string		true	"access_token"
//	@Param			node	body		PostNodeIn	true	"узел с сегментами"
//	@Success		200		{object}	PostNodeOut	"id узла"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/uzi/nodes [post]
func (h *Handler) PostNodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req PostNodeIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	segments := make([]*uzipb.CreateNodeIn_NestedSegment, 0, len(req.Segments))
	for _, v := range req.Segments {
		segments = append(segments, &uzipb.CreateNodeIn_NestedSegment{
			ImageId:   v.ImageID.String(),
			Contor:    v.Contor,
			Tirads_23: v.Tirads23,
			Tirads_4:  v.Tirads4,
			Tirads_5:  v.Tirads5,
		})
	}

	resp, err := h.adapter.UziAdapter.CreateNode(ctx, &uzipb.CreateNodeIn{
		UziId:     req.UziID.String(),
		Segments:  segments,
		Tirads_23: req.Tirads23,
		Tirads_4:  req.Tirads4,
		Tirads_5:  req.Tirads5,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetAllNodes получить все узлы узи
//
//	@Summary		получить все узлы узи
//	@Description	получить все узлы узи
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string		true	"access_token"
//	@Param			id		path		string					true	"uzi_id"
//	@Success		200		{object}	GetAllNodesOut	"id узла"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/uzi/uzis/{id}/nodes [get]
func (h *Handler) GetAllNodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	// TODO: в ответе пустые поля будут опущены, убрать теги omitempty.
	resp, err := h.adapter.UziAdapter.GetAllNodes(ctx, &uzipb.GetAllNodesIn{
		UziId: id,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// DeleteNode удалит узел
//
//	@Summary		удалит узел
//	@Description	удалит узел
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Success		200		{string}	string	"molodec"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/uzi/nodes/{id} [delete]
func (h *Handler) DeleteNode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.UziAdapter.DeleteNode(ctx, &uzipb.DeleteNodeIn{Id: id})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// PatchNode обновит узел
//
//	@Summary		обновит узел
//	@Description	обновит узел
//	@Tags			uzi
//	@Produce		json
//	@Param			node	body		PatchNodeIn		true	"узел с сегментами"
//	@Success		200		{object}	PatchNodeOut	"обновленный узел"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/uzi/nodes/{id} [patch]
func (h *Handler) PatchNode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var req PatchNodeIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	resp, err := h.adapter.UziAdapter.UpdateNode(ctx, &uzipb.UpdateNodeIn{
		Id:        id,
		Tirads_23: req.Tirads23,
		Tirads_4:  req.Tirads4,
		Tirads_5:  req.Tirads5,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp.Node); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// TODO: нет валидации что если ноды нет
// PostSegment добавит новый сегмент к указанному узлу
//
//	@Summary		добавит новый сегмент к указанному узлу
//	@Description	добавит новый сегмент к указанному узлу
//	@Tags			uzi
//	@Produce		json
//	@Param			node	body		PostSegmentIn	true	"сегмент"
//	@Success		200		{object}	PostSegmentOut	"id узла"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/uzi/segments [post]
func (h *Handler) PostSegment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req PostSegmentIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	resp, err := h.adapter.UziAdapter.CreateSegment(ctx, &uzipb.CreateSegmentIn{
		ImageId:   req.ImageID.String(),
		NodeId:    req.NodeID.String(),
		Contor:    req.Contor,
		Tirads_23: req.Tirads23,
		Tirads_4:  req.Tirads4,
		Tirads_5:  req.Tirads5,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// DeleteSegment удалит сегмент
//
//	@Summary		удалит сегмент
//	@Description	удалит сегмент, ЕСЛИ У УЗЛА НЕ ОСТАНЕТСЯ СЕГМЕНТОВ, ОН ТОЖЕ БУДЕТ УДАЛЕН
//	@Tags			uzi
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Success		200		{string}	string	"molodec"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/uzi/segments/{id} [delete]
func (h *Handler) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.UziAdapter.DeleteSegment(ctx, &uzipb.DeleteSegmentIn{Id: id})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// PatchSegment обновит сегмент
//
//	@Summary		обновит сегмент
//	@Description	обновит сегмент
//	@Tags			uzi
//	@Produce		json
//	@Param			node	body		PatchSegmentIn	true	"узел с сегментами"
//	@Success		200		{object}	PatchSegmentOut	"обновленный узел"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/uzi/segments/{id} [patch]
func (h *Handler) PatchSegment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var req PatchSegmentIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	resp, err := h.adapter.UziAdapter.UpdateSegment(ctx, &uzipb.UpdateSegmentIn{
		Id:        id,
		Tirads_23: req.Tirads23,
		Tirads_4:  req.Tirads4,
		Tirads_5:  req.Tirads5,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp.Segment); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}
