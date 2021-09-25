package recovery

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("[PANIC] %s panic recovered:\n%s\n",
					time.Now().Format(time.RFC3339), err)

				c.Abort()
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
