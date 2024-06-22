package middlewares

import (
	"net/http"

	"github.com/fadhlinw/clean-gin/lib"
	"github.com/fadhlinw/clean-gin/error"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type ErrorMiddleware struct {
	logger lib.Logger
}

func NewErrorMiddleware(
	logger lib.Logger,
) ErrorMiddleware {
	return ErrorMiddleware{
		logger: logger,
	}
}

func (m ErrorMiddleware) Setup() {}

func (m ErrorMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case error.Http:
				abortErrorResponse(c, e.Description, e.StatusCode)
			case *mysql.MySQLError:
				if e.Number == 1062 {
					abortErrorResponse(c, "Duplicate entry", http.StatusConflict)
				} else {
					abortErrorResponse(c, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			default:
				if e == gorm.ErrRecordNotFound {
					abortErrorResponse(c, "Record not found", http.StatusNotFound)
				} else {
					abortErrorResponse(c, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
		}
	}
}
