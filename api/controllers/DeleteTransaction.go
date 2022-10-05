package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/gorilla/mux"
)

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	transactionInstance := models.Transaction{}

	// get the slug with the name "id"
	params := mux.Vars(r)
	uuid := params["uuid"]

	// get the request body into the todo struct
	json.NewDecoder(r.Body).Decode(&transactionInstance)

	_, err := models.DeleteTransaction(&transactionInstance, uuid)

	// if an error is found, send it to the client
	if err != nil {
		fmt.Println("Error found", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
