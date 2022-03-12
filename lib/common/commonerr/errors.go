package commonerr

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"

	"github.com/firdasafridi/parkinglot/lib/common/log"
)

//
// An ErrorMessage represents an error message format and list of error.
//
type ErrorMessage struct {
	ErrorList []*ErrorFormat `json:"error_list"`
	Code      int            `json:"code"`
}

//
// An ErrorFormat represents an error message format and code that we used.
//
type ErrorFormat struct {
	ErrorName        string `json:"error_name"`
	ErrorDescription string `json:"error_description"`
}

// OrderErrorFormat Error Format for order
type OrderErrorFormat struct {
	Code   string `json:"code"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

//
// Common internal server error message
//
const (
	InternalServerName        = "internal_server_error"
	InternalServerDescription = "The server is unable to complete your request"
)

//
// Common error unautorized
//
const (
	UnauthorizedErrorName        = "access_denied"
	UnauthorizedErrorDescription = "Authorization failed by filter."
)

//
// Common bad request error message
//
var DefaultBadRequest = ErrorFormat{
	ErrorName:        "bad_request",
	ErrorDescription: "Your request resulted in error",
}

//
// Default Not Found error message
//
const (
	NotFound            = "not_found"
	NotFoundDescription = "Page not found"
)

//
// Create new error message
//
func NewErrorMessage() *ErrorMessage {
	return &ErrorMessage{}
}

//
// Set bad request
//
func (em *ErrorMessage) SetBadRequest() *ErrorMessage {
	em.Code = http.StatusBadRequest
	return em
}

//
// Set bad request with array
//
func NewBadRequestWithArray(uid string, errMessage []string) *ErrorMessage {
	errList := &ErrorMessage{
		Code:      http.StatusBadRequest,
		ErrorList: []*ErrorFormat{},
	}
	for _, message := range errMessage {
		errList.Append(uid, message)
	}
	return errList
}

//
// SetErrorValidator contains setter error from github.com/go-playground/validator
//
func (em *ErrorMessage) SetErrorValidator(err error) *ErrorMessage {
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return em
		}

		for _, err := range err.(validator.ValidationErrors) {
			em.Append(err.Field(), err.Tag())
		}

	}
	return em
}

//
// SetNewError is function return new error message.
// It support to set code, error name, and error description
//
func SetNewError(code int, errorName, errDesc string) *ErrorMessage {
	return &ErrorMessage{
		Code: code,
		ErrorList: []*ErrorFormat{
			{
				ErrorName:        errorName,
				ErrorDescription: errDesc,
			},
		},
	}
}

//
// SetDefaultNewNotFound returns a default 404 error page not found
//
func SetDefaultNewNotFound() *ErrorMessage {
	return SetNewError(404, NotFound, NotFoundDescription)
}

// SetNewNotFound returns a 404 Not Found error with customized messages.
func SetNewNotFound(errorName, errDesc string) *ErrorMessage {
	return SetNewError(http.StatusNotFound, errorName, errDesc)
}

//
// SetNewBadRequest is function return new error message with bad request standard code(400).
// It support to set error name and error description
//
func SetNewBadRequest(errorName, errDesc string) *ErrorMessage {
	return SetNewError(http.StatusBadRequest, errorName, errDesc)
}

//
// Set404 is function return new error message with not found standard code(404).
// It support to set error name and error description
//
func Set404() *ErrorMessage {
	return SetNewError(http.StatusNotFound, "404", "404 page not found")
}

//
// SetNewBadRequest is function return new error message with bad request standard code(400).
// It support to set error name and error description using error format
//
func SetNewBadRequestByFormat(ef *ErrorFormat) *ErrorMessage {
	return &ErrorMessage{
		Code: http.StatusBadRequest,
		ErrorList: []*ErrorFormat{
			ef,
		},
	}
}

//
// SetDefaultNewBadRequest returns default bad request error with http code 400
//
func SetDefaultNewBadRequest() *ErrorMessage {
	return SetNewBadRequestByFormat(&DefaultBadRequest)
}

//
// SetNewInternalError is function return new error message with internal server error standard code(500).
//
func SetNewInternalError() *ErrorMessage {
	return SetNewError(http.StatusInternalServerError, InternalServerName, InternalServerDescription)
}

//
// SetNewUnauthorizedError is function return new error message with unauthorized error code(401).
// It support to set error name and error description
//
func SetNewUnauthorizedError(errorName, errDesc string) *ErrorMessage {
	return SetNewError(http.StatusUnauthorized, errorName, errDesc)
}

//
// SetNewUnauthorizedError is function return new error message with unauthorized error code(401).
// It support to set error name and error description
//
func SetDefaultUnauthorized() *ErrorMessage {
	return SetNewUnauthorizedError(UnauthorizedErrorName, UnauthorizedErrorDescription)
}

//
// Append is function add error to existing error message.
// It support to set error name and error description.
//
func (errorMessage *ErrorMessage) Append(errorName, errDesc string) *ErrorMessage {
	errorMessage.ErrorList = append(errorMessage.ErrorList, &ErrorFormat{
		ErrorName:        errorName,
		ErrorDescription: errDesc,
	})
	return errorMessage
}

//
// AppendFormat is function add error to existing error message.
// It support to set error name and error description using error format
//
func (errorMessage *ErrorMessage) AppendFormat(ef *ErrorFormat) *ErrorMessage {
	errorMessage.ErrorList = append(errorMessage.ErrorList, ef)
	return errorMessage
}

//
// GetListError is function to get list error message.
//
func (errorMessage *ErrorMessage) GetListError() []*ErrorFormat {
	return errorMessage.ErrorList
}

//
// GetCode is function to get code.
//
func (errorMessage *ErrorMessage) GetCode() int {
	return errorMessage.Code
}

//
// Get error byte
//
func (errorMessage *ErrorMessage) Marshal() []byte {
	b, _ := json.Marshal(errorMessage)
	return b
}

//
// Get string
//
func (errorMessage *ErrorMessage) ToString() string {
	return string(errorMessage.Marshal())
}

// Error to implement error interface
func (errorMessage *ErrorMessage) Error() string {
	return errorMessage.ToString()
}

// Errorln print the error
func (errorMessage *ErrorMessage) Errorln(err ...interface{}) *ErrorMessage {
	log.Errorln(err...)
	return errorMessage
}

// Debugln print the error
func (errorMessage *ErrorMessage) Debugln(err ...interface{}) *ErrorMessage {
	log.Debugln(err...)
	return errorMessage
}
