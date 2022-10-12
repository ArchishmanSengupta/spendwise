package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ArchishmanSengupta/expense-tracker/utils"
)

func TestDeleteTransaction(t *testing.T) {
	ts := utils.TestServer{Server: Server}

	t.Run("Bad Request with non existing transaction", func(t *testing.T) {
		statusCode, _, _ := ts.Delete(t, "/transactions/8fe8aa41-29a6-401a-9cc4-98874f3472a3", "")
		if statusCode != http.StatusNotFound {
			t.Errorf("want %d status code; got %d", http.StatusNotFound, statusCode)
		}
	})
	t.Run("Valid request", func(t *testing.T) {
		statusCode, _, resBody := ts.Delete(t, "/transactions/8fe8aa41-29a6-401a-9cc4-98874f3472a2", "")

		if statusCode != http.StatusOK {
			t.Errorf("want %d status code; got %d", http.StatusOK, statusCode)
		} else {
			var response map[string]interface{}

			err := json.Unmarshal(resBody, &response)
			if err != nil {
				t.Fatal("Error unmarshalling response body: ", err.Error())
			}

		}
	})
}
