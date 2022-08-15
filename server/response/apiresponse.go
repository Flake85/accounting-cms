package response

import (
	"encoding/json"
	"net/http"
)

type ApiResponseBody struct {
	Data    interface{} `json:"data,omitempty"`
	Error   ApiError    `json:"error,omitempty"`
}

type ApiError struct {
	Message string 		`json:"message,omitempty"`
	Data 	interface{} `json:"data,omitempty"`
}

func NewOkResponse(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(ApiResponseBody{Data: data})
}

func NewErrorResponse(code int, message string, w http.ResponseWriter) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ApiResponseBody{
		Error: ApiError{Message: message},
	})
}

func NewErrorResponseWithData(code int, message string, data interface{}, w http.ResponseWriter) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ApiResponseBody{
		Error: ApiError{
			Message: message,
			Data: data,
		},
	})
}
