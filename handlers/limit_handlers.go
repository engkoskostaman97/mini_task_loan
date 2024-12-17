package handlers

import (
	"encoding/json"
	"net/http"

	"xyz_multifinance/dto"
	"xyz_multifinance/models"
	"xyz_multifinance/repository"
)

type Handler struct{}

func CreateLimit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request dto.CreateLimitRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Validate input values
	if request.ConsumerID <= 0 {
		http.Error(w, "Invalid consumer_id", http.StatusBadRequest)
		return
	}
	if request.Tenor <= 0 {
		http.Error(w, "Invalid tenor value", http.StatusBadRequest)
		return
	}
	if request.Amount <= 0 {
		http.Error(w, "Invalid amount value", http.StatusBadRequest)
		return
	}

	limit := models.Limit{
		ConsumerID: uint(request.ConsumerID),
		Tenor:      request.Tenor,
		Amount:     request.Amount,
	}

	if err := repository.CreateLimit(&limit); err != nil {
		http.Error(w, "Failed to create limit", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Limit created successfully",
		"data":    limit,
	}
	json.NewEncoder(w).Encode(response)
}
