package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// Get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}
