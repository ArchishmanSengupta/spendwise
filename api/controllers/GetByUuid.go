package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/gorilla/mux"
)

func GetByUuid(w http.ResponseWriter, r *http.Request) {
	transactionInstance := models.Transaction{}
	params := mux.Vars(r)
	uuid := params["uuid"]

	dbConn := cmd.DbConn

	transaction, err := transactionInstance.Retrieve(dbConn, map[string]interface{}{"uuid": uuid})

	if err != nil {
		fmt.Println("Error found in Controller of GetByUuid---->", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(transaction)
}
