/*********************************************************
 * THIS FILE IS THE CONTROLLER OF THE CREATE TRANSACTION *
 *********************************************************/

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/api/serializers"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	transactionInstance := models.Transaction{}
	json.NewDecoder(r.Body).Decode(&transactionInstance)

	var errors = make(map[string]string)

	if transactionInstance.Amount == 0 {
		errors["amount"] = "Amount is required, and must be greater than 0"
	}
	if transactionInstance.Type == "" {
		errors["type"] = "Type is required"
	}

	if len(errors) > 0 {
		utils.SendError(w, errors, http.StatusBadRequest)
		return
	}

	transaction, err := transactionInstance.Insert()

	if err != nil {
		utils.SendError(w, nil, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	transactionSerializer := serializers.TransactionSerializer{
		Transactions: []*models.Transaction{transaction},
		Many:         false,
	}

	//send the created transaction to the response
	_ = json.NewEncoder(w).Encode(transactionSerializer.Serialize()["data"])

}
