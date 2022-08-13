package response

type ApiResponse struct {
	Code int
	Body ApiResponseBody
}

type ApiResponseBody struct {
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type ApiError struct {
	Message string 		`json:"message,omitempty"`
	Data 	interface{} `json:"data,omitempty"`
}

func NewOkResponse(data interface{}) ApiResponse {
	return ApiResponse{
		Body: ApiResponseBody{
			Data: data,
		},
	}
}

func NewErrorResponse(code int, message string) ApiResponse {
	return ApiResponse{
		Code: code,
		Body: ApiResponseBody{
			Error: ApiError{
				Message: message,
			},
		},
	}
}

func NewErrorResponseWithData(code int, message string, data interface{}) ApiResponse {
	return ApiResponse{
		Code: code,
		Body: ApiResponseBody{
			Error: ApiError{
				Message: message,
				Data: data,
			},
		},
	}
}