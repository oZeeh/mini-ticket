// middlewares/auth.go
package middlewares

import (
	customErrors "backend/custom_errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.Error(customErrors.ErrForbidden)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, customErrors.ErrForbidden
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.Error(customErrors.ErrForbidden)
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
		if err != nil {
			c.Error(customErrors.ErrForbidden)
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Set("role", claims["role"])
		c.Next()
	}
}
