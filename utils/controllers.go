package utils

import "net/http"

// SendError - send the error response to the client
// @Params - w (http.ResponseWriter), err (error), status (int)
func SendError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)

	var errMessage string

	if status == http.StatusNotFound {
		errMessage = "The specified resource was not found"
	} else if status == http.StatusBadRequest {
		errMessage = "One or more parameters are missing"
	} else {
		errMessage = "Something went wrong"
	}

	w.Write([]byte(errMessage))
}
