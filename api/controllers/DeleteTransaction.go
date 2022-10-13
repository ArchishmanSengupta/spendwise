/*******************************************************
 * THIS FILE IS THE CONTROLLER OF THE DELETE OPERATION *
 *******************************************************/
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
	"github.com/gorilla/mux"
)

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	transactionInstance := models.Transaction{}

	// get the slug with the name "id"
	params := mux.Vars(r)
	uuid := params["uuid"]

	// Get the DBconn from the main.go
	dbConn := cmd.DbConn

	err := transactionInstance.Delete(dbConn, map[string]interface{}{"uuid": uuid})

	if err != nil {
		switch err {
		case utils.ErrResourceNotFound:
			utils.SendError(w, nil, http.StatusNotFound)
		default:
			utils.LogError(r, err.Error())
			utils.SendError(w, nil, http.StatusInternalServerError)
		}
		return
	}

	response := map[string]interface{}{
		"uuid":   uuid,
		"object": "transaction",
		"delete": true,
	}

	_ = json.NewEncoder(w).Encode(response)
}
