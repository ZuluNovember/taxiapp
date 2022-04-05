package routers

import (
	_ "github.com/ZuluNovember/taxiapp/driver/docs"
	"github.com/ZuluNovember/taxiapp/driver/routers/api"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	driver := r.Group("/driver")
	{
		driver.GET("/import", api.ImportCsv)
		driver.POST("/add", api.AddDriver)
		driver.POST("/match", api.FindDriver)
	}
	return r
}
