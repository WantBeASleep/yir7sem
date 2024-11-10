package med

import (
	"context"
	"net/http"
	"strconv"
	"yir/gateway/internal/custom"
	"yir/gateway/internal/entity/workermodel"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

func (c *MedController) PostWorker(w http.ResponseWriter, r *http.Request) {
	worker := workermodel.AddMedicalWorkerRequest{}
	validate := validator.New()
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&worker); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	ctx := context.Background()
	data, err := c.Service.AddMedWorker(ctx, &worker)
	if err != nil {
		http.Error(w, "Failed to add worker. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *MedController) PutWorker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	worker := workermodel.MedicalWorkerUpdateRequest{}
	validate := validator.New()
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&worker); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	ctx := context.Background()
	data, err := c.Service.UpdateMedWorker(ctx, id, &worker)
	if err != nil {
		http.Error(w, "Failed to add worker. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *MedController) GetWorkerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	data, err := c.Service.GetMedWorkerByID(ctx, id)
	if err != nil {
		http.Error(w, "Failed to add worker. Please try again later.", http.StatusBadGateway)
		return
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (c *MedController) GetWorkersList(w http.ResponseWriter, r *http.Request) {
	var limit, offset uint64 = 0, 0
	ctx := context.Background()
	data, err := c.Service.GetMedWorkers(ctx, limit, offset)
	if err != nil {
		http.Error(w, "Failed to add worker. Please try again later.", http.StatusBadGateway)
		return
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (c *MedController) GetWorkerPatients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	data, err := c.Service.GetPatientsByMedWorker(ctx, id)
	if err != nil {
		http.Error(w, "Failed to add worker. Please try again later.", http.StatusBadGateway)
		return
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}
