package uzi

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	brokeradapters "yir/gateway/internal/adapters/broker"
	grpcadapters "yir/gateway/internal/adapters/grpc"
	"yir/gateway/internal/domain"

	// uziuploadpb "yir/gateway/internal/generated/broker/produce/uziupload"
	uzipb "yir/gateway/internal/generated/grpc/client/uzi"
	"yir/gateway/internal/repository"
)

type Handler struct {
	grpcadapter   grpcadapters.Adapter
	brokeradapter brokeradapters.Adapter
	dao           repository.DAO
}

func New(
	grpcadapter grpcadapters.Adapter,
	brokeradapter brokeradapters.Adapter,
	dao repository.DAO,
) *Handler {
	return &Handler{
		grpcadapter:   grpcadapter,
		brokeradapter: brokeradapter,
		dao:           dao,
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

	uziResp, err := h.grpcadapter.UziAdapter.CreateUzi(ctx, &uzipb.CreateUziIn{
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
	// if err := h.brokeradapter.SendUziUpload(&uziuploadpb.UziUpload{UziId: uziResp.Id}); err != nil {
	// 	http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
	// 	return
	// }

	w.WriteHeader(200)
}
