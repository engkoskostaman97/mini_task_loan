package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"xyz_multifinance/database"
	"xyz_multifinance/dto"
	"xyz_multifinance/models"
	"xyz_multifinance/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate = validator.New()

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("Response encoding error: %v\n", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

func parseDate(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", date)
}

// CREATE 
func CreateConsumer(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateConsumerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	if err := validate.Struct(req); err != nil {
		log.Printf("Validation error: %v\n", err)
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Validation error: " + err.Error()})
		return
	}

	parsedDate, err := parseDate(req.DateOfBirth)
	if err != nil {
		log.Printf("Error parsing date_of_birth: %v\n", err)
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid date_of_birth format. Expected YYYY-MM-DD."})
		return
	}

	consumer := models.Consumer{
		NIK:          req.NIK,
		FullName:     req.FullName,
		LegalName:    req.LegalName,
		PlaceOfBirth: req.PlaceOfBirth,
		DateOfBirth:  parsedDate,
		Salary:       req.Salary,
		KTPPhoto:     req.KTPPhoto,
		SelfiePhoto:  req.SelfiePhoto,
	}

	if err := repository.CreateConsumer(&consumer); err != nil {
		log.Printf("Database error: %v\n", err)
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create consumer"})
		return
	}

	writeJSONResponse(w, http.StatusCreated, consumer)
}

// READ ALL Consumers
func GetAllConsumers(w http.ResponseWriter, r *http.Request) {
	var consumers []models.Consumer

	db := database.DB

	if err := db.Preload("Limits").Preload("Transactions").Find(&consumers).Error; err != nil {
		log.Printf("Database error: %v\n", err)
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch consumers"})
		return
	}

	writeJSONResponse(w, http.StatusOK, consumers)
}

// READ Consumer by ID
func GetConsumerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Error parsing ID: %v\n", err)
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
		return
	}

	var consumer models.Consumer

	db := database.DB

	if err := db.Preload("Limits").Preload("Transaction").First(&consumer, id).Error; err != nil {
		log.Printf("Database error: %v\n", err)
		writeJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Consumer not found"})
		return
	}

	writeJSONResponse(w, http.StatusOK, consumer)
}

// DELETE Consumer
func DeleteConsumer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Error parsing ID: %v\n", err)
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
		return
	}

	if err := repository.DeleteConsumer(uint(id)); err != nil {
		log.Printf("Database error: %v\n", err)
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete consumer"})
		return
	}

	writeJSONResponse(w, http.StatusNoContent, nil)
}
