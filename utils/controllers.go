package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Type    string      `json:"type"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

// SendError - send the error response to the client
// @Params - w (http.ResponseWriter), err (error), status (int)
func SendError(w http.ResponseWriter, err interface{}, status int) {
	errorResponse := ErrorResponse{}
	errorResponse.Type = "invalid_request_error"
	errorResponse.Code = status

	if status == http.StatusNotFound {
		errorResponse.Message = "The specified resource was not found"
	} else if status == http.StatusBadRequest {
		errorResponse.Message = "One or more parameters are missing"
		errorResponse.Errors = err
	} else {
		errorResponse.Message = "Something went wrong"
	}

	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorResponse)
}
