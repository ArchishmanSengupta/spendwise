/*******************************************************
 * THIS FILE IS THE CONTROLLER OF THE DELETE OPERATION *
 *******************************************************/
package controllers

import (
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
	"github.com/gorilla/mux"
)

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
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
			utils.SendError(w, err, http.StatusNotFound)
		default:
			utils.SendError(w, err, http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)

	// send response
	w.Write([]byte(http.StatusText(http.StatusNoContent)))
}
