package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/api/serializers"
	"github.com/ArchishmanSengupta/expense-tracker/cmd"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactionInstance := models.Transaction{}

	typeFromTheUrl := r.URL.Query().Get("type")     //check if type is empty
	DateFromTheUrl := r.URL.Query().Get("date")     //check if type is empty
	amountFromTheUrl := r.URL.Query().Get("amount") //check if type is empty

	attributeMap := make(map[string]interface{}, 0)

	var filter bool = false

	if typeFromTheUrl != "" {
		attributeMap["type"] = typeFromTheUrl
		filter = true
	}

	if amountFromTheUrl != "" {
		attributeMap["amount"] = amountFromTheUrl
		filter = true
	}

	if DateFromTheUrl != "" {
		currentTime := time.Now().String()
		DateFromTheUrl = currentTime[:10]
		fmt.Println("DateFromTheUrl--->", DateFromTheUrl)
		attributeMap["date"] = DateFromTheUrl
		filter = true
	}
	var transactions []*models.Transaction
	var err error
	dbConn := cmd.DbConn
	if filter {
		transactions, err = transactionInstance.Filter(dbConn, attributeMap)
	} else {
		transactions, err = transactionInstance.GetAllTransactions(typeFromTheUrl, amountFromTheUrl)
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
