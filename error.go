package request

import (
	"fmt"
	"net/http"
)

//ResponseError Error type for http status code not 200
type ResponseError struct {
	StatusCode int
	Request    *http.Request
	Message    string
}

//NewResponseError New ResponseError
func NewResponseError(statusCode int, request *http.Request, message string) *ResponseError {
	if statusCode == 200 {
		panic("statusCode 200 but ResponseError?")
	}
	return &ResponseError{statusCode, request, message}
}

func (r ResponseError) Error() string {
	e := "Status Code "
	e += fmt.Sprint(r.StatusCode)
	e += " in " + r.Request.Method + " Request"
	e += " to " + r.Request.URL.String() + "\n"
	e += r.Message
	return e
}
