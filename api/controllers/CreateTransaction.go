/*********************************************************
 * THIS FILE IS THE CONTROLLER OF THE CREATE TRANSACTION *
 *********************************************************/

package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/api/serializers"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	transactionInstance := models.Transaction{}
	json.NewDecoder(r.Body).Decode(&transactionInstance)

	if transactionInstance.Amount == 0 || transactionInstance.Type == "" || transactionInstance.Amount < 0 {
		utils.SendError(w, errors.New("Missing Fields"), http.StatusBadRequest)
		return
	}
	transaction, err := transactionInstance.Insert()

	if err != nil {
		if err == utils.ErrResourceNotFound {
			utils.SendError(w, err, http.StatusNotFound)
		} else {
			utils.SendError(w, err, http.StatusInternalServerError)
		}
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
