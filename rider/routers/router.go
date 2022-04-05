package routers

import (
	_ "github.com/ZuluNovember/taxiapp/rider/docs"
	"github.com/ZuluNovember/taxiapp/rider/middleware/jwt"
	"github.com/ZuluNovember/taxiapp/rider/routers/api"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.Authenticate)
	driver := r.Group("/rider")
	driver.Use(jwt.JWT())
	{
		driver.POST("/match", api.Match)
	}
	return r
}
