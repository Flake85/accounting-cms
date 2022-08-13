package response

type ApiResponse struct {
	Code int
	Body ApiResponseBody
}

type ApiResponseBody struct {
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type BaseMessage struct {
	Message string `json:"message"`
}

type ApiError struct {
	BaseMessage
	Data interface{} `json:"data,omitempty"`
}

func NewOkResponse(data interface{}) ApiResponse {
	return ApiResponse{
		Body: ApiResponseBody{
			Data: data,
		},
	}
}

func NewErrorResponse(code int, message BaseMessage, data interface{}) ApiResponse {
	return ApiResponse{
		Code: code,
		Body: ApiResponseBody{
			Error: ApiError{
				BaseMessage: message,
				Data: data,
			},
		},
	}
}

func NewBaseMessage(message string) BaseMessage {
	return BaseMessage{Message: message}
}