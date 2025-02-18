package logger

import (
	"github.com/gin-gonic/gin"
)

// Logger
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// logger logic goes here
	}
}
