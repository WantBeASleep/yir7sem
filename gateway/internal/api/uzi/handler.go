package uzi

// TODO: большая проблема: то что рисуем на выход в сваггер != тому что туда реально уходит (уходит GRPC)

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

	w.WriteHeader(200)
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

	w.WriteHeader(200)
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

	w.WriteHeader(200)
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
//	@Router			/uzi/uzi/{id} [get]
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

	w.WriteHeader(200)
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

	w.WriteHeader(200)
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

	w.WriteHeader(200)
}
