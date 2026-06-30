package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	domainerrors "github.com/ascend/phoenix-multichannel-mkp/pkg/errors"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorBody  `json:"error,omitempty"`
}

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PaginatedResponse struct {
	Items interface{} `json:"items"`
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Success: true, Data: data})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Response{Success: true, Data: data})
}

func Fail(c *gin.Context, err error) {
	var domErr *domainerrors.DomainError
	if errors.As(err, &domErr) {
		statusCode := domainErrorToHTTP(domErr.Err)
		c.JSON(statusCode, Response{
			Success: false,
			Error:   &ErrorBody{Code: domErr.Code, Message: domErr.Message},
		})
		return
	}
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Error:   &ErrorBody{Code: "INTERNAL_ERROR", Message: "an unexpected error occurred"},
	})
}

func domainErrorToHTTP(err error) int {
	switch {
	case errors.Is(err, domainerrors.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, domainerrors.ErrAlreadyExists):
		return http.StatusConflict
	case errors.Is(err, domainerrors.ErrInvalidInput):
		return http.StatusBadRequest
	case errors.Is(err, domainerrors.ErrUnauthorized):
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
