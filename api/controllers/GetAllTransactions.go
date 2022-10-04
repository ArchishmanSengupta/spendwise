package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transaction := models.Transaction{}

	// get all transactions
	transactions, err := models.GetAllTransactions(&transaction)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(transactions)
}
