/*****************************************************
 * THIS FILE IS THE CONTROLLER OF GET ALL TRANSACTIONS *
 *****************************************************/
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/api/serializers"
	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	transactionInstance := models.Transaction{}
	limit, offset := utils.ExtractPaginationParams(r)

	typeFromTheUrl := r.URL.Query().Get("type")
	dateFromTheUrl := r.URL.Query().Get("date")
	amountFromTheUrl := r.URL.Query().Get("amount")

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
		transactions, err = transactionInstance.Filter(dbConn, attributeMap, limit, offset)
	} else {
		transactions, err = transactionInstance.GetAllTransactions(typeFromTheUrl, dateFromTheUrl, limit, offset, cmd.DbConn)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Transaction data Serialization
	transactionSerializer := serializers.TransactionSerializer{
		Transactions: transactions,
		Many:         true,
	}
	// Get next pagination URL
	nextPaginationUrl := utils.GetPaginationUrl(limit, offset, models.TotalTransactionCount, r)

	hasMore := true
	if nextPaginationUrl == "" {
		hasMore = false
	}
	responseMap := map[string]interface{}{
		"next":     nextPaginationUrl,
		"count":    models.TotalTransactionCount,
		"has_more": hasMore,
		"data":     transactionSerializer.Serialize()["data"],
	}
	//send the created transaction to the response
	_ = json.NewEncoder(w).Encode(responseMap)
}
