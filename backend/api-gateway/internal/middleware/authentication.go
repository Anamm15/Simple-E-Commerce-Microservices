package middleware

import (
	"net/http"
	"strings"

	helper "api-gateway/internal/helper/jwt"
	"api-gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := util.BuildResponseFailed("Failed ", "Access token not found", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenStr := strings.Split(authHeader, " ")[1]
		if tokenStr == "" {
			response := util.BuildResponseFailed("Failed ", "Access token not found", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims, err := helper.ValidateJWT(tokenStr)
		if err != nil {
			response := util.BuildResponseFailed("Failed ", "Invalid access token", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set("token", tokenStr)
		ctx.Set("user_id", claims.UserID)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}
