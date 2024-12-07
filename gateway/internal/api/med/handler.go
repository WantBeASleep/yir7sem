package med

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WantBeASleep/goooool/gtclib"

	adapters "gateway/internal/adapters"
	pb "gateway/internal/generated/grpc/client/med"

	"github.com/gorilla/mux"
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
//	@Tags			med
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Success		200		{object}	GetDoctorOut
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/med/doctors [get]
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

// GetDoctorPatients Получить пациентов врача
//
//	@Summary		Получить пациентов врача
//	@Description	Получить пациентов врача
//	@Tags			med
//	@Produce		json
//	@Param			token	header		string					true	"access_token"
//	@Success		200		{object}	GetDoctorPatientsOut	"пациенты"
//	@Failure		500		{string}	string					"Internal Server Error"
//	@Router			/med/doctors/patients [get]
func (h *Handler) GetDoctorPatients(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	XUserID := r.Header.Get("x-user_id")

	res, err := h.adapter.MedAdapter.GetDoctorPatients(ctx, &pb.GetDoctorPatientsIn{DoctorId: XUserID})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(res.Patients); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	w.WriteHeader(200)
}

// UpdateDoctor обновляет поля у доктора
//
//	@Summary		Обновить врача
//	@Description	Обновить врача
//	@Tags			med
//	@Produce		json
//	@Param			token	header		string			true	"access_token"
//	@Param			body	body		UpdateDoctorIn	true	"обновляемые значения"
//	@Success		200		{object}	UpdateDoctorOut	"врач"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/med/doctors [patch]
func (h *Handler) UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	XUserID := r.Header.Get("x-user_id")

	var req UpdateDoctorIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	res, err := h.adapter.MedAdapter.UpdateDoctor(ctx, &pb.UpdateDoctorIn{
		Id:   XUserID,
		Org:  req.Org,
		Job:  req.Job,
		Desc: req.Desc,
	})
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

// PostPatient создать пациента
//
//	@Summary		создать пациента
//	@Description	создать пациента
//	@Tags			med
//	@Produce		json
//	@Param			token	header		string			true	"access_token"
//	@Param			body	body		PostPatientIn	true	"данные пациента"
//	@Success		200		{object}	PostPatientOut	"id пациента"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/med/patient [post]
func (h *Handler) PostPatient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req PostPatientIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	res, err := h.adapter.MedAdapter.CreatePatient(ctx, &pb.CreatePatientIn{
		Fullname:   req.FullName,
		Email:      req.Email,
		Policy:     req.Policy,
		Active:     req.Active,
		Malignancy: req.Malignancy,
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

// GetPatient Получить пациента
//
//	@Summary		Получить пациента
//	@Description	Получить пациента
//	@Tags			med
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Param			id		path		string	true	"patient_id"
//	@Success		200		{object}	GetPatientOut
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/med/patient/{id} [get]
func (h *Handler) GetPatient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	res, err := h.adapter.MedAdapter.GetPatient(ctx, &pb.GetPatientIn{Id: id})
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

// UpdatePatient обновляет поля пациента
//
//	@Summary		обновляет поля пациента
//	@Description	обновляет поля пациента
//	@Tags			med
//	@Produce		json
//	@Param			token	header		string				true	"access_token"
//	@Param			id		path		string				true	"patient_id"
//	@Param			body	body		UpdatePatientIn		true	"обновляемые значения"
//	@Success		200		{object}	UpdatePatientOut	"пациент"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/med/patient/{id} [patch]
func (h *Handler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	XUserID := r.Header.Get("x-user_id")
	id := mux.Vars(r)["id"]

	var req UpdatePatientIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	res, err := h.adapter.MedAdapter.UpdatePatient(ctx, &pb.UpdatePatientIn{
		DoctorId:    XUserID,
		Id:          id,
		Active:      req.Active,
		Malignancy:  req.Malignancy,
		LastUziDate: gtclib.Timestamp.TimePointerTo(req.LastUziDate),
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(res.Patient); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	w.WriteHeader(200)
}

// PostCard создать карту
//
//	@Summary		создать карту
//	@Description	создать карту
//	@Tags			med
//	@Produce		json
//	@Param			token	header		string		true	"access_token"
//	@Param			body	body		PostCardIn	true	"данные карты"
//	@Success		200		{object}	PostCardOut	"id карты"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/med/card [post]
func (h *Handler) PostCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	XUserID := r.Header.Get("x-user_id")

	var req PostCardIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	res, err := h.adapter.MedAdapter.CreateCard(ctx, &pb.CreateCardIn{
		Card: &pb.Card{
			DoctorId:  XUserID,
			PatientId: req.PatientID.String(),
			Diagnosis: req.Diagnosis,
		},
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

// GetCard Получить карту
//
//	@Summary		Получить карту
//	@Description	Получить карту
//	@Tags			med
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Param			id		path		string	true	"patient_id"
//	@Success		200		{object}	GetCardOut
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/med/card/{id} [get]
func (h *Handler) GetCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	XUserID := r.Header.Get("x-user_id")
	patientID := mux.Vars(r)["id"]

	res, err := h.adapter.MedAdapter.GetCard(ctx, &pb.GetCardIn{
		DoctorId:  XUserID,
		PatientId: patientID,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(res.Card); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	w.WriteHeader(200)
}

// UpdateCard обновить карту
//
//	@Summary		обновить карту
//	@Description	обновить карту
//	@Tags			med
//	@Produce		json
//	@Param			token	header		string			true	"access_token"
//	@Param			id		path		string			true	"patient_id"
//	@Param			body	body		UpdateCardIn	true	"обновляемые значения"
//	@Success		200		{object}	UpdateCardOut	"карта"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/med/card/{id} [patch]
func (h *Handler) UpdateCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	XUserID := r.Header.Get("x-user_id")
	id := mux.Vars(r)["id"]

	var req UpdateCardIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	res, err := h.adapter.MedAdapter.UpdateCard(ctx, &pb.UpdateCardIn{
		Card: &pb.Card{
			DoctorId:  XUserID,
			PatientId: id,
			Diagnosis: req.Diagnosis,
		},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(res.Card); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	w.WriteHeader(200)
}
