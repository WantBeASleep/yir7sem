package download

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"gateway/internal/repository"

	"github.com/gorilla/mux"
)

type Handler struct {
	dao repository.DAO
}

func New(
	dao repository.DAO,
) *Handler {
	return &Handler{
		dao: dao,
	}
}

// GetUzi Получение узи
//
//	@Summary		Получение узи
//	@Description	Получение узи
//	@Tags			download
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Param			uzi_id	path		string	true	"id узи"
//	@Success		200		{file}		File	"Изображение УЗИ"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/download/uzi/{id} [get]
func (h *Handler) GetUzi(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	file, err := h.dao.NewFileRepo().GetFile(ctx, filepath.Join(id, id))
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	// TODO: переписать dao, Добавить content-type
	w.Header().Set("Content-Type", "image/tiff")
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, fmt.Sprintf("не удалось вернуть изображение: %v", err), 500)
		return
	}
}

// GetImage Получение image uzi
//
//	@Summary		Получение image uzi
//	@Description	Получение image uzi
//	@Tags			download
//	@Produce		json
//	@Param			token		header		string	true	"access_token"
//	@Param			uzi_id		path		string	true	"id узи"
//	@Param			image_id	path		string	true	"id image"
//	@Success		200			{file}		File	"Изображение кадра Узи"
//	@Failure		500			{string}	string	"Internal Server Error"
//	@Router			/download/uzi/{uzi_id}/image/{image_id} [get]
func (h *Handler) GetImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uziID := mux.Vars(r)["uzi_id"]
	imageID := mux.Vars(r)["image_id"]

	file, err := h.dao.NewFileRepo().GetFile(
		ctx,
		filepath.Join(
			uziID,
			imageID,
			imageID,
		),
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	// TODO: переписать dao, Добавить content-type
	w.Header().Set("Content-Type", "image/png")
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, fmt.Sprintf("не удалось вернуть изображение: %v", err), 500)
		return
	}
}
