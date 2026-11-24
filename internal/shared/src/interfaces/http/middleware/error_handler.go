package middleware

import (
	"errors"
	apperror "go-wallet/internal/shared/src/domain/error"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	err := c.Errors.Last().Err

	var appErr *apperror.AppError

	if errors.As(err, &appErr) {
		status := mapHTTPStatus(appErr.Kind)

		if status >= http.StatusInternalServerError {
			slog.Error("application error", "error", err)
		} else {
			slog.Info("application error", "error", err)
		}

		c.AbortWithStatusJSON(status, gin.H{
			"error": gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
			},
		})
		return
	}

	slog.Error("unexpected error", "error", err)

	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": gin.H{
			"code":    "INTERNAL_ERROR",
			"message": "Internal server error",
		},
	})
}

func mapHTTPStatus(kind apperror.Kind) int {
	switch kind {
	case apperror.NotFound:
		return http.StatusNotFound
	case apperror.Conflict:
		return http.StatusConflict
	case apperror.Validation:
		return http.StatusBadRequest
	case apperror.Unauthorized:
		return http.StatusUnauthorized
	case apperror.Forbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
