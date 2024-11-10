package med

import (
	"context"
	"net/http"
	"strconv"
	"yir/gateway/internal/entity/cardmodel"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
)

func (c *MedController) PostCard(w http.ResponseWriter, r *http.Request) {
	patientInfo := cardmodel.PatientInformation{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&patientInfo); err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&patientInfo); err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	if err := c.Service.PostCard(ctx, &patientInfo); err != nil {
		http.Error(w, "Failed to add card. Please try again later.", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *MedController) PutCard(w http.ResponseWriter, r *http.Request) {
	card := cardmodel.PatientCard{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&card); err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	resp, err := c.Service.PutCard(ctx, &card)
	if err != nil {
		http.Error(w, "Failed to add card. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *MedController) GetCards(w http.ResponseWriter, r *http.Request) {
	var limit, offset uint64 = 0, 0 // по идее jsonка должна быть
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	ctx := context.Background()
	resp, err := c.Service.GetCards(ctx, limit, offset)
	if err != nil {
		http.Error(w, "Failed to get cards. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *MedController) GetCardByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)

	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	resp, err := c.Service.GetCardByID(ctx, id)
	if err != nil {
		http.Error(w, "Failed to get card by id. Please try again later.", http.StatusBadGateway)
		return
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *MedController) DeleteCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)

	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	if err := c.Service.DeleteCard(ctx, id); err != nil {
		http.Error(w, "Failed to get card by id. Please try again later.", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}
