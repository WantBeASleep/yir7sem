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

// PostCard godoc
// @Summary      Add Patient Card
// @Description  Creates a new patient card with information about the patient, card details, and the medical worker.
// @Tags         cards
// @Accept       json
// @Produce      json
// @Param        body  body  cardmodel.PatientInformation  true  "Patient Information"
// @Success      200   {string}   string  "Card created successfully"
// @Failure      400   {string}   string  "Invalid request payload"
// @Failure      502   {string}   string  "Failed to add card"
// @Router       /med/card/add [post]
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

// PutCard godoc
// @Summary      Update Patient Card
// @Description  Updates information of an existing patient card.
// @Tags         cards
// @Accept       json
// @Produce      json
// @Param        body  body  cardmodel.PatientCard  true  "Patient Card Data"
// @Success      200   {object}   cardmodel.PatientCard  "Updated card information"
// @Failure      400   {string}   string  "Invalid request payload"
// @Failure      502   {string}   string  "Failed to update card"
// @Router       /med/card/update [put]
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

// GetCards godoc
// @Summary      Get Patient Cards
// @Description  Retrieves a list of patient cards with pagination options.
// @Tags         cards
// @Accept       json
// @Produce      json
// @Param        limit   query  int  false  "Limit"  default(10)
// @Param        offset  query  int  false  "Offset"  default(0)
// @Success      200   {object}   cardmodel.PatientCardList  "List of patient cards"
// @Failure      502   {string}   string  "Failed to get cards"
// @Router       /med/card/list [get]
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

// GetCardByID godoc
// @Summary      Get Patient Card by ID
// @Description  Retrieves details of a specific patient card by its ID.
// @Tags         cards
// @Accept       json
// @Produce      json
// @Param        id   path   int  true  "Card ID"
// @Success      200   {object}   cardmodel.PatientCard  "Patient card details"
// @Failure      400   {string}   string  "Invalid ID format"
// @Failure      502   {string}   string  "Failed to get card by ID"
// @Router       /med/card/{id} [get]
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

// DeleteCard godoc
// @Summary      Delete Patient Card
// @Description  Deletes a specific patient card by its ID.
// @Tags         cards
// @Accept       json
// @Produce      json
// @Param        id   path   int  true  "Card ID"
// @Success      200   {string}   string  "Card deleted successfully"
// @Failure      400   {string}   string  "Invalid ID format"
// @Failure      502   {string}   string  "Failed to delete card"
// @Router       /med/cards/delete/{id} [delete]
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
