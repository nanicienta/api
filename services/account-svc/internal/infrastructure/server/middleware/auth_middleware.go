// Package middleware provides middleware functions for the HTTP server
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nanicienta/api/account-svc/internal/auth"
	"strings"
)

// Auth is a middleware function that checks for a valid JWT token in the Authorization header
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.AbortWithStatusJSON(
				401, gin.H{
					"message": "Unauthorized",
				},
			)
			context.Abort()
			return
		}
		claims, err := auth.ValidateToken(
			strings.Replace(
				tokenString,
				"Bearer ",
				"",
				1,
			),
		)
		if err != nil {
			context.AbortWithStatusJSON(
				401, gin.H{
					"message": "Unauthorized",
				},
			)
			context.Abort()
			return
		}
		context.Set("kimosUserId", claims.ID)
		context.Next()
	}
}
