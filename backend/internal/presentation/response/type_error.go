package response

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/response/constant"
)

func NewConflictError(err error) *ResponseError {
	return NewResponseError(
		err,
		http.StatusConflict,
		fmt.Sprintf(
			constant.ConflictErrorMessage,
			err.Error(),
		),
	)
}

func NewUnauthorizedError(err error) *ResponseError {
	return NewResponseError(
		err,
		http.StatusUnauthorized,
		constant.UnauthorizedErrorMessage,
	)
}

func NewNotFoundError(err error) *ResponseError {
	return NewResponseError(
		err,
		http.StatusNotFound,
		fmt.Sprintf(
			constant.NotFoundErrorMessage,
			err.Error(),
		),
	)
}

func NewInternalServerError(err error) *ResponseError {
	return NewResponseError(
		err,
		http.StatusInternalServerError,
		constant.InternalServerErrorMessage,
	)
}

func NewBadRequestError(err error) *ResponseError {
	return NewResponseError(
		err,
		http.StatusBadRequest,
		fmt.Sprintf(
			constant.BadRequestErrorMessage,
			err.Error(),
		),
	)
}

func NewTimeoutError() *ResponseError {
	return NewResponseError(
		errors.New(constant.RequestTimeoutErrorMessage),
		http.StatusRequestTimeout,
		constant.RequestTimeoutErrorMessage,
	)
}

func NewForbiddenError(err error) *ResponseError {
	return NewResponseError(
		err,
		http.StatusForbidden,
		fmt.Sprintf(
			constant.ForbiddenErrorMessage,
		),
	)
}
