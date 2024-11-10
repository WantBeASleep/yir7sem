package med

import (
	"context"
	"net/http"
	"strconv"
	"yir/gateway/internal/custom"
	"yir/gateway/internal/entity/patientmodel"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

func (c *MedController) PostPatient(w http.ResponseWriter, r *http.Request) {
	Info := patientmodel.PatientInformation{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&Info); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(Info); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	ctx := context.Background()
	if err := c.Service.AddPatient(ctx, &Info); err != nil {
		http.Error(w, "Failed to add patient. Please try again later.", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *MedController) PutPatient(w http.ResponseWriter, r *http.Request) {
	Info := patientmodel.PatientInformation{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&Info); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(Info); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	ctx := context.Background()
	if err := c.Service.UpdatePatient(ctx, &Info); err != nil {
		http.Error(w, "Failed to add patient. Please try again later.", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *MedController) GetPatientInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	patient, err := c.Service.GetPatientInfoByID(ctx, id)
	if err != nil {
		http.Error(w, "Failed to get patient by id. Please try again later.", http.StatusBadGateway)
		return
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.NewEncoder(w).Encode(&patient); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (c *MedController) GetPatientList(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	data, err := c.Service.GetPatientList(ctx)
	if err != nil {
		http.Error(w, "Failed to get patient list. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (c *MedController) GetPatientsShots(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "NOT IMPLEMENTED", http.StatusBadGateway)
}
