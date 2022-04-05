package main

import (
	"log"
	"net/http"

	"github.com/ZuluNovember/taxiapp/driver/models"
	"github.com/ZuluNovember/taxiapp/driver/pkg/settings"
	"github.com/ZuluNovember/taxiapp/driver/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	settings.Setup()
	models.Setup()
}

// @title           Driver API
// @version         1.0
// @description     Sample driver location API

// @contact.name   Eniz Coban
// @contact.email  zulunovember.c@gmail.com

// @host      localhost:8000
// @BasePath  /
func main() {
	gin.SetMode(settings.Configuration.Server.RunMode)

	routersInit := routers.InitRouter()
	endPoint := settings.Configuration.Server.AcceptAddress + ":" + settings.Configuration.Server.EndPoint

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
