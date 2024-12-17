package routes

import (
	"xyz_multifinance/handlers"
	"github.com/gorilla/mux"
)

// InitializeRoutes sets up all the API routes
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// Consumer routes
	router.HandleFunc("/consumers", handlers.CreateConsumer).Methods("POST")
	router.HandleFunc("/consumers/{id}", handlers.GetConsumerByID).Methods("GET")
	router.HandleFunc("/consumers/{id}", handlers.DeleteConsumer).Methods("DELETE")
	router.HandleFunc("/consumers", handlers.GetAllConsumers).Methods("GET")

	// Limit routes
	router.HandleFunc("/limit", handlers.CreateLimit).Methods("POST")

  // Transaction routes 
	router.HandleFunc("/transaction", handlers.CreateTransaction).Methods("POST")
	router.HandleFunc("/transaction/{id}", handlers.GetTransactionByID).Methods("GET")
	router.HandleFunc("/transaction/{id}", handlers.UpdateTransaction).Methods("PUT")
	router.HandleFunc("/transaction/{id}", handlers.DeleteTransaction).Methods("DELETE")
	router.HandleFunc("/transaction", handlers.GetAllTransactions).Methods("GET")

	return router
}
