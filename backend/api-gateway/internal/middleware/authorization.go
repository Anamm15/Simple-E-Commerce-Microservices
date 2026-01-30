package middleware

import (
	"net/http"
	"slices"

	"api-gateway/internal/helper/enum"
	"api-gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

func AuthorizeRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleInterface, exists := c.Get("role")
		if !exists {
			res := util.BuildResponseFailed("Failed ", "Role not found", nil)
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		}

		userRole, ok := roleInterface.(enum.AccountRole)
		if !ok {
			roleStr, okStr := roleInterface.(string)
			if !okStr {
				res := util.BuildResponseFailed("Failed", "Internal Server Error: Invalid Role Type", nil)
				c.AbortWithStatusJSON(http.StatusInternalServerError, res)
				return
			}
			userRole = enum.AccountRole(roleStr)
		}

		if slices.Contains(allowedRoles, string(userRole)) {
			c.Next()
			return
		}

		res := util.BuildResponseFailed("Failed ", "Unauthorized", nil)
		c.JSON(http.StatusForbidden, res)
		c.Abort()
	}
}
