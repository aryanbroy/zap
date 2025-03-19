package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Successful bool   `json:"successful"`
	Status     int    `json:"status"`
	Error      string `json:"error"`
}

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error, status int) *ErrorResponse {
	return &ErrorResponse{
		Successful: false,
		Error:      err.Error(),
		Status:     status,
	}
}

func CustomError(err string, status int) *ErrorResponse {
	return &ErrorResponse{
		Successful: false,
		Error:      err,
		Status:     status,
	}
}
