package e

import (
	"fmt"
	"net/http"
)

type Error interface {
	Code() int
	Detail() string
}

type HttpError struct {
	detail string
	code   int
}

func NewHttpError(code int, detail string) HttpError {
	return HttpError{
		code:   code,
		detail: detail,
	}
}

func (e HttpError) Error() string {
	return fmt.Sprintf(`code: %d, detail: '%s'`, e.code, e.detail)
}

func (e HttpError) Code() int {
	return e.code
}

func (e HttpError) Detail() string {
	return e.detail
}

func NewInternal(detail string) Error {
	return HttpError{
		detail: detail,
		code:   http.StatusInternalServerError,
	}
}

func NewInternalf(template string, args ...interface{}) Error {
	return HttpError{
		detail: fmt.Sprintf(template, args...),
		code:   http.StatusInternalServerError,
	}
}

func NewNotFound(detail string) Error {
	return HttpError{
		detail: detail,
		code:   http.StatusNotFound,
	}
}

func NewBadRequest(detail string) Error {
	return HttpError{
		detail: detail,
		code:   http.StatusBadRequest,
	}
}
