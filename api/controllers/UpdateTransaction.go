/****************************************************************
 * THIS FILE IS THE CONTROLLER FOR THE UPDATING THE TRANSACTION *
 ****************************************************************/
package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/api/serializers"
	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
	"github.com/gorilla/mux"
)

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	transactionInstance := models.Transaction{}

	// get the slug with the name "id"
	params := mux.Vars(r)
	idStr := params["id"]

	//Convert String id to int
	id, _ := strconv.Atoi(idStr)

	if transactionInstance.Amount == 0 || transactionInstance.Amount < 0 || transactionInstance.Type == "" {
		utils.SendError(w, errors.New("Missing fields"), http.StatusBadRequest)
		return
	}

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
	fmt.Println("Transaction Serialization---->", transactionSerializer.Serialize())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseMap := map[string]interface{}{
		"status_code": "success",
		"message":     "Transaction updated successfully",
		"data":        transactionSerializer.Serialize()["data"],
	}
	_ = json.NewEncoder(w).Encode(responseMap)
}
