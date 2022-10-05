package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	transactionInstance := models.Transaction{}

	json.NewDecoder(r.Body).Decode(&transactionInstance)

	transactionInstance.CreatedAt = time.Now()
	transactionInstance.UpdatedAt = time.Now()

	transaction, err := models.CreateTransaction(&transactionInstance)

	//Error Handling
	if err != nil {
		fmt.Println("Error found", err)
	}
	//content type
	w.Header().Set("Content-Type", "application/json")

	//encode
	json.NewEncoder(w).Encode(transaction)

}
