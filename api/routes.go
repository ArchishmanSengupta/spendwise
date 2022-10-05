package api

import (
	"github.com/ArchishmanSengupta/expense-tracker/api/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {

	// User Routes
	r.HandleFunc("/transactions", controllers.GetAllTransactions).Methods("GET")
	r.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions/{uuid}", controllers.GetByUuid).Methods("GET")
	r.HandleFunc("/transactions/{uuid}", controllers.UpdateTransaction).Methods("PUT")
	r.HandleFunc("/transactions/{uuid}", controllers.DeleteTransaction).Methods("DELETE")
}
