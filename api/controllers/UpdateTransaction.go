package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/gorilla/mux"
)

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	transactionInstance := models.Transaction{}

	// get the slug with the name "id"
	params := mux.Vars(r)
	idStr := params["id"]

	// convert id from string to int
	id, _ := strconv.Atoi(idStr)

	// get the todo with this id first
	_, err := transactionInstance.Retrieve(id)

	// if an error is found, send it to the client
	if err != nil {
		fmt.Println("Error found", err)
	}

	// get the request body into the todo struct
	json.NewDecoder(r.Body).Decode(&transactionInstance)

	todo, err := transactionInstance.UpdateTransaction(id)

	// if an error is found, send it to the client
	if err != nil {
		fmt.Println("Error found", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
