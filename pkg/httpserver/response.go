// Package httpserver holds generic Gin response helpers so handlers write a
// consistent success/error envelope. It depends only on pkg/apperror and knows
// nothing about this app's domain — sentinel-to-AppError mapping stays in the
// caller (e.g. internal/apperr).
package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/okdev/marketplace-sync/pkg/apperror"
)

type GetSuccessResponse struct {
	Code    string `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

// NewSuccessResponse writes data as the raw 200 body.
func NewSuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// NewSuccessResponseWithData wraps data in the standard success envelope.
func NewSuccessResponseWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, GetSuccessResponse{
		Code:    "0",
		Message: "Success",
		Data:    data,
	})
}

// NewErrorResponse writes err as a JSON error. An AppError is sent with its own
// code + HTTP status; any other error falls back to a 500 carrying the error
// text as the message.
func NewErrorResponse(c *gin.Context, err error) {
	if e, ok := apperror.IsAppError(err); ok {
		c.JSON(e.HTTPStatusCode, e)
		return
	}
	c.JSON(http.StatusInternalServerError, apperror.AppError{
		Code:    apperror.ErrInternalCode,
		Message: apperror.ErrorMessage(err.Error()),
	})
}
