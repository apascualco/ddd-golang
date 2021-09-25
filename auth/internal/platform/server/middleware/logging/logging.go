package logging

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		if c.Request.URL.RawQuery != "" {
			path = path + "?" + c.Request.URL.RawQuery
		}
		c.Next()

		timestamp := time.Now()
		latency := timestamp.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		fmt.Printf("[HTTP] %v | %3d | %13v | %15s | %-7s %#v\n",
			timestamp.Format(time.RFC3339),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
