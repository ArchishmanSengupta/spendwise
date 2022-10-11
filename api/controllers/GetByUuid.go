/*******************************************************************************************
 *     THIS FILE IS THE CONTROLLER FOR GETTING TRANSACTION RESPONSE BASED ON THE UUID.     *
 * THIS CONTROLLER WILL BE CALLED WHEN THE USER HITS THE ENDPOINT /API/TRANSACTIONS/{UUID} *
 *******************************************************************************************/
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/api/serializers"
	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
	"github.com/gorilla/mux"
)

func GetByUuid(w http.ResponseWriter, r *http.Request) {
	transactionInstance := models.Transaction{}
	params := mux.Vars(r)
	uuid := params["uuid"]

	fmt.Println("uuid", uuid)
	dbConn := cmd.DbConn

	transaction, err := transactionInstance.Retrieve(dbConn, map[string]interface{}{"uuid": uuid})

	fmt.Println("Get By UUID: ", transaction)

	// if an error is found, send it to the client
	if err != nil {
		switch err {
		case utils.ErrResourceNotFound:
			utils.SendError(w, err, http.StatusNotFound)
		default:
			utils.SendError(w, err, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	transactionSerializer := serializers.TransactionSerializer{
		Transactions: []*models.Transaction{transaction},
		Many:         true,
	}

	// send the todo to the client
	_ = json.NewEncoder(w).Encode(transactionSerializer.Serialize()["data"])
}
