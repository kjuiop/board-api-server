package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"runtime"
)

func RecoveryErrorReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 1024)
				runtime.Stack(buf, false)
				errMsg := fmt.Sprintf("recovered from panic : %s", buf)
				slog.Error(errMsg)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
