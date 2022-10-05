package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/gorilla/mux"
)

func GetByUuid(w http.ResponseWriter, r *http.Request) {
	transactionInstance := models.Transaction{}
	params := mux.Vars(r)
	uuid := params["uuid"]

	// get all transactions
	// transactions, err := models.GetByUuid(&transaction, uuid)

	transaction, err := transactionInstance.Retrieve(uuid)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(transaction)
}
