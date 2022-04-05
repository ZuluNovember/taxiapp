package main

import (
	"log"
	"net/http"

	"github.com/ZuluNovember/taxiapp/rider/pkg/settings"
	"github.com/ZuluNovember/taxiapp/rider/routers"
	"github.com/ZuluNovember/taxiapp/rider/util"
	"github.com/gin-gonic/gin"
)

func init() {
	settings.Setup()
	util.Setup()
}

// @title           Rider API
// @version         1.0
// @description     Sample matching API

// @contact.name   Eniz Coban
// @contact.email  zulunovember.c@gmail.com

// @host      localhost:8080
// @BasePath  /
func main() {
	gin.SetMode(settings.Configuration.Server.RunMode)

	routersInit := routers.InitRouter()
	endPoint := ":" + settings.Configuration.Server.EndPoint

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
