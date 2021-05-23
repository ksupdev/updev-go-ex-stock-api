package interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GeneralInterceptor1 - call this methods to add interceptor
func GenerateInterceptor1(c *gin.Context) {
	token := c.Query("token")
	if token == "1234" {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"errpr": "invalid token"})
		c.Abort()
	}

}
