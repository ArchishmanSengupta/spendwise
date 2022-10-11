package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jmoiron/sqlx"
)

type TestServer struct {
	*httptest.Server
}

func (ts *TestServer) Post(t *testing.T, urlPath string, reqBody string, contentType string) (int, http.Header, []byte) {
	req, err := http.NewRequest("POST", ts.URL+urlPath, bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}
func ClearTestDatabase(db *sqlx.DB) {
	_, _ = db.Exec(`TRUNCATE TABLE todo	;`)
}
func SeedTransactions(db *sqlx.DB) error {
	query := `INSERT INTO transaction (id, amount, type) VALUES 
			(1, 23422, 'credit'),
			(2, 42233, 'debit')
			`
	_, err := db.Exec(query)

	return err
}
