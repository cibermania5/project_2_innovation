package operations

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type clientOptions *options.ClientOptions
var clOpt clientOptions


type (
	// Response is the http json response schema
	Response struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Content interface{} `json:"content"`
	}
)

// NewResponse is the Response struct factory function.
func NewResponse(status int, message string, content interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Content: content,
	}
}

func ResponseWriter(res http.ResponseWriter, statusCode int, message string, data interface{}) error {
	res.WriteHeader(statusCode)
	httpResponse := NewResponse(statusCode, message, data)
	err := json.NewEncoder(res).Encode(httpResponse)
	return err
}