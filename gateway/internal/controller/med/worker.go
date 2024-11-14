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

// PostWorker godoc
// @Summary      Add New Worker
// @Description  Creates a new medical worker record with details like name, organization, and job.
// @Tags         workers
// @Accept       json
// @Produce      json
// @Param        body  body  workermodel.AddMedicalWorkerRequest  true  "Worker Information"
// @Success      200   {object}  workermodel.MedicalWorker        "Created worker data"
// @Failure      400   {string}   string  "Invalid request payload"
// @Failure      502   {string}   string  "Failed to add worker"
// @Router       /med/worker/add [post]
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

// PutWorker godoc
// @Summary      Update Worker Information
// @Description  Updates details of an existing medical worker.
// @Tags         workers
// @Accept       json
// @Produce      json
// @Param        id    path      int                                 true  "Worker ID"
// @Param        body  body      workermodel.MedicalWorkerUpdateRequest  true  "Updated Worker Information"
// @Success      200   {object}  workermodel.MedicalWorker           "Updated worker data"
// @Failure      400   {string}   string  "Invalid request payload"
// @Failure      502   {string}   string  "Failed to update worker"
// @Router       /med/workers/update/{id} [put]
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

// GetWorkerID godoc
// @Summary      Get Worker by ID
// @Description  Retrieves information of a specific worker by ID.
// @Tags         workers
// @Accept       json
// @Produce      json
// @Param        id    path      int                             true  "Worker ID"
// @Success      200   {object}  workermodel.MedicalWorker       "Worker information"
// @Failure      400   {string}   string  "Invalid ID format"
// @Failure      502   {string}   string  "Failed to get worker information"
// @Router       /med/worker/{id} [get]
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

// GetWorkersList godoc
// @Summary      List All Workers
// @Description  Retrieves a list of all medical workers.
// @Tags         workers
// @Accept       json
// @Produce      json
// @Success      200   {object}  workermodel.MedicalWorkerList    "List of workers"
// @Failure      502   {string}   string  "Failed to get workers list"
// @Router       /med/worker/list [get]
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

// GetWorkerPatients godoc
// @Summary      Get Patients by Worker ID
// @Description  Retrieves a list of patients associated with a specific medical worker by ID.
// @Tags         workers
// @Accept       json
// @Produce      json
// @Param        id    path      int                                  true  "Worker ID"
// @Success      200   {object}  workermodel.MedicalWorkerWithPatients "Worker with patients"
// @Failure      400   {string}   string  "Invalid ID format"
// @Failure      502   {string}   string  "Failed to get patients for worker"
// @Router       /med/worker/patients/{medWorkerId} [get]
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
