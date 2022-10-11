package tests

import (
	"net/http"
	"testing"

	"github.com/ArchishmanSengupta/expense-tracker/utils"
)

func TestCreateTransaction(t *testing.T) {
	ts := utils.TestServer{Server: Server}

	//When the request body does not contain an amount

	t.Run("Bad Request without required fields", func(t *testing.T) {
		reqBody := `{
			"type": "credit"
		}`
		statusCode, _, _ := ts.Post(t, "/transactions", reqBody, "")
		if statusCode != http.StatusBadRequest {
			t.Errorf("Want %d status code; got %d", http.StatusBadRequest, statusCode)
		}
	})
	//When the request body does not contain a type

	t.Run("Bad Request without required fields", func(t *testing.T) {
		reqBody := `{
			"amount": 23422
		}`
		statusCode, _, _ := ts.Post(t, "/transactions", reqBody, "")
		if statusCode != http.StatusBadRequest {
			t.Errorf("Want %d status code; got %d", http.StatusBadRequest, statusCode)
		}
	})
	// When the request is valid with all fields
	t.Run("Valid request", func(t *testing.T) {
		reqBody := `{
			"amount": 23421 ,
			"type": "credit"
		}`

		statusCode, _, _ := ts.Post(t, "/transactions", reqBody, "")

		if statusCode != http.StatusOK {
			t.Errorf("want %d status code; got %d", http.StatusOK, statusCode)
		}
	})
}
