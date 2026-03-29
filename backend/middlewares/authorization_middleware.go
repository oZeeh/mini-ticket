// middlewares/authorization.go
package middlewares

import (
	customErrors "backend/custom_errors"
	"backend/users/enums"

	"github.com/gin-gonic/gin"
)

func RequireRole(roles ...enums.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			c.Error(customErrors.ErrForbidden)
			c.Abort()
			return
		}

		role := enums.Role(roleVal.(string))

		for _, r := range roles {
			if role == r {
				c.Next()
				return
			}
		}

		c.Error(customErrors.ErrForbidden)
		c.Abort()
	}
}
