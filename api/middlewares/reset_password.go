package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/fadhlinw/clean-gin/constants"
	"github.com/fadhlinw/clean-gin/domains"
	"github.com/fadhlinw/clean-gin/lib"
)

// JWTResetPasswordMiddleware middleware for jwt reset password authentication
type JWTResetPasswordMiddleware struct {
	service domains.AuthService
	logger  lib.Logger
}

// NewJWTResetPasswordMiddleware creates new jwt reset password middleware
func NewJWTResetPasswordMiddleware(service domains.AuthService, logger lib.Logger) JWTResetPasswordMiddleware {
	return JWTResetPasswordMiddleware{
		service: service,
		logger:  logger,
	}
}

// Setup sets up jwt reset password middleware
func (m JWTResetPasswordMiddleware) Setup() {}

// Handler returns a Gin middleware function that handles JWT reset password authentication.
func (m JWTResetPasswordMiddleware) Handler() gin.HandlerFunc {
	m.logger.Info("JWT reset password middleware")
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := m.service.Authorize(authToken, constants.TypeResetToken)
			if authorized != nil {
				c.Request.Header.Set("user_id", authorized.UserID)
				c.Request.Header.Set("reset_token", authToken)
				c.Next()
				return
			}

			m.logger.Error(err)
			abortErrorResponse(c, err.Error(), http.StatusUnauthorized)
			return
		}
		abortErrorResponse(c, "invalid token", http.StatusUnauthorized)
	}
}