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

// PostPatient godoc
// @Summary      Add New Patient
// @Description  Creates a new patient record, including personal information and patient card details.
// @Tags         patients
// @Accept       json
// @Produce      json
// @Param        body  body  patientmodel.PatientInformation  true  "Patient Information"
// @Success      200   {string}   string  "Patient created successfully"
// @Failure      400   {string}   string  "Invalid request payload"
// @Failure      502   {string}   string  "Failed to add patient"
// @Router       /med/patients/create [post]
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

// PutPatient godoc
// @Summary      Update Patient Information
// @Description  Updates existing patient details.
// @Tags         patients
// @Accept       json
// @Produce      json
// @Param        body  body  patientmodel.PatientInformation  true  "Updated Patient Information"
// @Success      200   {string}   string  "Patient updated successfully"
// @Failure      400   {string}   string  "Invalid request payload"
// @Failure      502   {string}   string  "Failed to update patient"
// @Router       /med/patient/update/{id} [put]
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

// GetPatientInfo godoc
// @Summary      Get Patient Information by ID
// @Description  Retrieves detailed information of a specific patient by ID.
// @Tags         patients
// @Accept       json
// @Produce      json
// @Param        id   path   int  true  "Patient ID"
// @Success      200   {object}   patientmodel.PatientInformation  "Patient information"
// @Failure      400   {string}   string  "Invalid ID format"
// @Failure      502   {string}   string  "Failed to get patient information"
// @Router       /med/patient/info/{id} [get]
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

// GetPatientList godoc
// @Summary      List All Patients
// @Description  Retrieves a list of all registered patients.
// @Tags         patients
// @Accept       json
// @Produce      json
// @Success      200   {array}    patientmodel.Patient  "List of patients"
// @Failure      502   {string}   string  "Failed to get patient list"
// @Router       /med/patient/list [get]
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

// FIXME: Надо оделать в пациентах
// GetPatientsShots godoc
// @Summary      Retrieve Patients' Shots Information
// @Description  This endpoint is not yet implemented.
// @Tags         patients
// @Accept       json
// @Produce      json
// @Failure      502   {string}   string  "NOT IMPLEMENTED"
// @Router       /med/patients/shots/{id} [get]
func (c *MedController) GetPatientsShots(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "NOT IMPLEMENTED", http.StatusBadGateway)
}
