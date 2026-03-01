package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	HttpCode int
	Code     string
	Message  string
}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.Code, e.Message)
}

func NewError(httpCode int, code string, message string) Error {
	return Error{
		HttpCode: httpCode,
		Code:     code,
		Message:  message,
	}
}

func IsCustomError(target error) bool {
	_, ok := target.(Error)
	if !ok {
		return false
	}
	return true
}

func Unwrap(target error) Error {
	e, ok := target.(Error)
	if !ok {
		return ErrGeneral
	}
	return e
}

var (
	ErrGeneral                  = NewError(http.StatusInternalServerError, "01", "General Error")
	ErrBadRequest               = NewError(http.StatusBadRequest, "02", "Invalid request payload")
	ErrNotFound                 = NewError(http.StatusNotFound, "03", "Data not found")
	ErrLoanClosed               = NewError(http.StatusNotAcceptable, "04", "Loan is already closed")
	ErrNoOutstandingInstallment = NewError(http.StatusNotAcceptable, "05", "No Outstanding Installment")
	ErrInvalidAmount            = NewError(http.StatusBadRequest, "06", "Invalid Amount")
)
