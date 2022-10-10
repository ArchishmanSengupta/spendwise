package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/api/serializers"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	transactionInstance := models.Transaction{}

	json.NewDecoder(r.Body).Decode(&transactionInstance)

	if transactionInstance.Amount < 0 {
		fmt.Println("Amount cannot be negative")
		return
	}
	transaction, err := transactionInstance.Insert()

	//Error Handling
	if err != nil {
		fmt.Println("Error found in Controller of CreateTransaction---->", err)
	}
	//content type
	w.Header().Set("Content-Type", "application/json")

	// Transaction data Serialization
	transactionSerializer := serializers.TransactionSerializer{
		Transactions: []*models.Transaction{transaction},
		Many:         false,
	}

	//send the created transaction to the response
	_ = json.NewEncoder(w).Encode(transactionSerializer.Serialize()["data"])

}
