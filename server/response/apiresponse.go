package response

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Code int			 `json:"code"`
	Body ApiResponseBody `json:"body"`
}

type ApiResponseBody struct {
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type ApiError struct {
	Message string 		`json:"message,omitempty"`
	Data 	interface{} `json:"data,omitempty"`
}

func NewOkResponse(data interface{}, w http.ResponseWriter) {
	res := ApiResponse{
		Code: 200,
		Body: ApiResponseBody{
			Data: data,
		},
	}
	json.NewEncoder(w).Encode(res)
}

func NewErrorResponse(code int, message string, w http.ResponseWriter) {
	res := ApiResponse{
		Code: code,
		Body: ApiResponseBody{
			Error: ApiError{
				Message: message,
			},
		},
	}
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

func NewErrorResponseWithData(code int, message string, data interface{}, w http.ResponseWriter) {
	res := ApiResponse{
		Code: code,
		Body: ApiResponseBody{
			Error: ApiError{
				Message: message,
				Data: data,
			},
		},
	}
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}