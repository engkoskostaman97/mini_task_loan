package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"xyz_multifinance/database"
	"xyz_multifinance/dto"
	"xyz_multifinance/models"
	"xyz_multifinance/repository"
)

// Fungsi untuk memeriksa apakah nomor kontrak sudah ada
func contractNoExists(contractNo string) bool {
	var transaction models.Transaction
	err := database.DB.Where("contract_no = ?", contractNo).First(&transaction).Error
	return err == nil
}

// CreateTransaction 
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transactionDTO dto.CreateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&transactionDTO); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Log the received data to check if contract_no is populated
	fmt.Printf("Received transaction: %+v\n", transactionDTO)

	// Validate contract_no
	if transactionDTO.ContractNo == "" {
		http.Error(w, "Contract number is required", http.StatusBadRequest)
		return
	}

	if contractNoExists(transactionDTO.ContractNo) {
		http.Error(w, "Contract number already exists", http.StatusBadRequest)
		return
	}

	transaction := models.Transaction{
		ConsumerID:  uint(transactionDTO.ConsumerID),
		ContractNo:  transactionDTO.ContractNo,
		OTR:         transactionDTO.OTR,
		AdminFee:    transactionDTO.AdminFee,
		Installment: transactionDTO.Installment,
		Interest:    transactionDTO.Interest,
		AssetName:   transactionDTO.AssetName,
	}

	if err := repository.CreateTransaction(&transaction); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create transaction: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

// GetTransactionByID
func GetTransactionByID(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Path[len("/transaction/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var transaction models.Transaction
	if err := repository.GetTransactionByID(uint(id), &transaction); err != nil {
		http.Error(w, fmt.Sprintf("Transaction not found: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

// GetAllTransactions 
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transaction
	if err := repository.GetAllTransactions(&transactions); err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve transactions: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

// UpdateTransaction handles updating a transaction
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Validate transaction ID
	if transaction.ID == 0 {
		http.Error(w, "Transaction ID is required", http.StatusBadRequest)
		return
	}

	var existingTransaction models.Transaction
	if err := repository.GetTransactionByID(transaction.ID, &existingTransaction); err != nil {
		http.Error(w, fmt.Sprintf("Transaction not found: %v", err), http.StatusNotFound)
		return
	}

	if err := repository.UpdateTransaction(&transaction); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update transaction: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

// DeleteTransaction handles deleting a transaction by ID
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Path[len("/transaction/"):]

	// Validate the ID format
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var transaction models.Transaction
	if err := repository.GetTransactionByID(uint(id), &transaction); err != nil {
		http.Error(w, fmt.Sprintf("Transaction not found: %v", err), http.StatusNotFound)
		return
	}

	if err := repository.DeleteTransaction(uint(id)); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete transaction: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Transaction deleted successfully"}); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
