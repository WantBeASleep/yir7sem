package med

import (
	"encoding/json"
	"fmt"
	"net/http"

	adapters "gateway/internal/adapters"
	pb "gateway/internal/generated/grpc/client/med"
)

type Handler struct {
	adapter adapters.Adapter
}

func New(
	adapter adapters.Adapter,
) *Handler {
	return &Handler{
		adapter: adapter,
	}
}

// GetDoctor возвращает информацию о враче
//
//	@Summary		Получить информацию о враче
//	@Description	Получает информацию о враче
//	@Tags			doctors
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Success		200		{object}	med.Doctor
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/med/doctors [post]
func (h *Handler) GetDoctor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	XUserID := r.Header.Get("x-user_id")

	res, err := h.adapter.MedAdapter.GetDoctor(ctx, &pb.GetDoctorIn{Id: XUserID})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(res.Doctor); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	w.WriteHeader(200)
}
