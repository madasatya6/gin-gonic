package csrf

import (
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
)

func New() gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	})
}
