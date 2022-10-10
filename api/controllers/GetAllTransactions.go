package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/api/serializers"
	"github.com/ArchishmanSengupta/expense-tracker/cmd"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactionInstance := models.Transaction{}

	typeFromTheUrl := r.URL.Query().Get("type")     //check if type is empty
	dateFromTheUrl := r.URL.Query().Get("date")     //check if type is empty
	amountFromTheUrl := r.URL.Query().Get("amount") //check if type is empty

	attributeMap := make(map[string]interface{}, 0)

	var filter bool = false

	if typeFromTheUrl != "" {
		filter = true
	}

	if amountFromTheUrl != "" {
		filter = true
	}

	if dateFromTheUrl != "" {
		filter = true
	}
	attributeMap["type"] = typeFromTheUrl
	attributeMap["date"] = dateFromTheUrl
	attributeMap["amount"] = amountFromTheUrl

	var transactions []*models.Transaction
	var err error
	dbConn := cmd.DbConn
	if filter {
		transactions, err = transactionInstance.Filter(dbConn, attributeMap)
	} else {
		transactions, err = transactionInstance.GetAllTransactions(typeFromTheUrl, dateFromTheUrl)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")

	// Transaction data Serialization
	transactionSerializer := serializers.TransactionSerializer{
		Transactions: transactions,
		Many:         true,
	}

	//send the created transaction to the response
	_ = json.NewEncoder(w).Encode(transactionSerializer.Serialize()["data"])
}
