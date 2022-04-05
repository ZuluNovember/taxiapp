package jwt

import (
	"net/http"

	"github.com/ZuluNovember/taxiapp/rider/pkg/app"
	"github.com/ZuluNovember/taxiapp/rider/util"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		token := c.Query("token")
		if token == "" {
			appG.Response(http.StatusUnauthorized, "Unauthorized request", nil)
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			appG.Response(http.StatusUnauthorized, nil, err)
			c.Abort()
			return
		}
		if !claims.Authenticated {
			appG.Response(http.StatusUnauthorized, nil, err)
			c.Abort()
			return
		}
		c.Next()
	}
}
