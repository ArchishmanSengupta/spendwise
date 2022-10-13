package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
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
func LogError(r *http.Request, errMessage string) {

	// Log error in command line
	log.Printf("Error: %s\n", errMessage)
}

// ExtractPaginationParams: function to extract limit and offset from the url
func ExtractPaginationParams(r *http.Request) (int, int) {
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "5"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 5
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	return limit, offset
}

// GetPaginationUrl: generates the next pagination url
// Params - limit (int), offset (int), totalRows (int), r (*http.Request)
// Returns the nextPaginationUrl (string)
func GetPaginationUrl(limit, offset, totalRows int, r *http.Request) string {
	remainingRows := totalRows - offset

	if remainingRows <= limit {
		limit = remainingRows
	}

	offset = offset + limit

	var nextPaginationUrl string

	if offset != totalRows {
		parsedUrl, err := url.Parse(r.URL.String())
		if err != nil {
			return ""
		}

		queryParams := parsedUrl.Query()
		queryParams.Set("limit", fmt.Sprintf("%d", limit))
		queryParams.Set("offset", fmt.Sprintf("%d", offset))
		parsedUrl.RawQuery = queryParams.Encode()
		nextPaginationUrl = parsedUrl.RequestURI()
	}

	return nextPaginationUrl
}
