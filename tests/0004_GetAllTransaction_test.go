package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ArchishmanSengupta/expense-tracker/utils"
)

func TestGetAllTransactions(t *testing.T) {
	ts := utils.TestServer{Server: Server}

	t.Run("Check status code and response length", func(t *testing.T) {
		statusCode, _, resBody := ts.Get(t, "/transactions", "")

		if statusCode != http.StatusOK {
			t.Errorf("want %d status code; got %d", http.StatusOK, statusCode)
		} else {
			var response map[string]interface{}

			err := json.Unmarshal(resBody, &response)
			if err != nil {
				t.Fatal("Error unmarshalling response body: ", err.Error())
			}

			data := response["data"].([]interface{})

			if len(data) != 3 {
				t.Errorf("want %d as response data length; got %d", 3, len(data))
			}
		}
	})
}
