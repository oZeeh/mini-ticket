package middlewares

import (
	customErrors "backend/custom_errors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		switch {
		case errors.Is(err, customErrors.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case errors.Is(err, customErrors.ErrConflict):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case errors.Is(err, customErrors.ErrForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
	}
}
