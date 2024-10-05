package httpErrors

import (
	"errors"
	"fmt"
	"net/http"
)

type Http struct {
	Err        string            `json:"error,omitempty"`
	Message    string            `json:"message,omitempty"`
	Metadata   *HttpErrorOptions `json:"metadata,omitempty"`
	StatusCode int               `json:"statusCode"`
}

type HttpErrorOptions struct {
	Message string `json:"message,omitempty"`
}

func NewHttpError(statusCode int, err error, options *HttpErrorOptions) Http {
	return Http{
		Err:        err.Error(),
		Message:    http.StatusText(statusCode),
		Metadata:   options,
		StatusCode: statusCode,
	}
}

func (err Http) Error() string {
	return fmt.Sprintf("statusCode: %v, error: %v", err.StatusCode, err.Err)
}

func NotFoundError(err error, httpOptions *HttpErrorOptions) Http {
	return NewHttpError(http.StatusNotFound, err, httpOptions)
}

func BadRequestError(err error, httpOptions *HttpErrorOptions) Http {
	return NewHttpError(http.StatusBadRequest, err, httpOptions)
}

func UnauthorizedError(err error, httpOptions *HttpErrorOptions) Http {
	return NewHttpError(http.StatusUnauthorized, err, httpOptions)
}

func InternalServerError(err error, httpOptions *HttpErrorOptions) Http {
	return NewHttpError(http.StatusInternalServerError, err, httpOptions)
}

func ForbiddenError(httpOptions *HttpErrorOptions) Http {
	return NewHttpError(http.StatusForbidden, errors.New("Forbidden"), httpOptions)
}
