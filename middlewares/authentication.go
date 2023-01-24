package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "secret" {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
}
