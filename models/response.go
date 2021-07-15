package models

type Response struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(status int, message string, data interface{}) *Response {
	return &Response{status, message, data}
}
