/****************************************************************
 * THIS FILE IS THE CONTROLLER FOR THE UPDATING THE TRANSACTION *
 ****************************************************************/
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/api/serializers"
	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
	"github.com/gorilla/mux"
)

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	transactionInstance := models.Transaction{}

	// get the slug with the name "id"
	params := mux.Vars(r)
	idStr := params["id"]

	//Convert String id to int
	id, _ := strconv.Atoi(idStr)

	//get transaction with this id first
	transactionRecord, err := transactionInstance.Retrieve(cmd.DbConn, map[string]interface{}{"id": id})

	// if an error is found, send it to the client
	if err != nil {
		switch err {
		case utils.ErrResourceNotFound:
			utils.SendError(w, err, http.StatusNotFound)
			return
		default:
			utils.SendError(w, err, http.StatusInternalServerError)
			return
		}
	}
	// get the request body into the transaction struct
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
	transactionRecord.Amount = transactionInstance.Amount
	transactionRecord.Type = transactionInstance.Type
	_, err = transactionRecord.UpdateTransaction(cmd.DbConn, map[string]interface{}{"id": id})

	// if an error is found, send it to the client
	if err != nil {
		switch err {
		case utils.ErrResourceNotFound:
			utils.SendError(w, err, http.StatusNotFound)
			return
		default:
			utils.SendError(w, err, http.StatusInternalServerError)
			return
		}
	}
	transactionSerializer := serializers.TransactionSerializer{
		Transactions: []*models.Transaction{transactionRecord},
		Many:         false,
	}

	responseMap := map[string]interface{}{
		"status_code": "success",
		"message":     "Transaction updated successfully",
		"data":        transactionSerializer.Serialize()["data"],
	}
	_ = json.NewEncoder(w).Encode(responseMap)
}
