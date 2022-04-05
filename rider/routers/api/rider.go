package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZuluNovember/taxiapp/rider/pkg/app"
	"github.com/ZuluNovember/taxiapp/rider/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Location struct {
	Lon float64 `json:"lon" validate:"gte=-180.0,lte=180.0,required"`
	Lat float64 `json:"lat" validate:"gte=-90.0,lte=90.0,required"`
}

// @Summary Authenticate user and sends a jwt as response
// @Tags Auth
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func Authenticate(c *gin.Context) {
	appG := app.Gin{C: c}

	token, err := util.GenerateToken()
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil, err)
		return
	}

	appG.Response(http.StatusOK, gin.H{"token": token}, nil)
}

// @Summary Find a driver close to given coordinates
// @Tags rider
// @Accept json
// @Produce  json
// @Param coordinate body api.Location true "Coordinate"
// @Param token query string true "Jwt string"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /rider/match [post]
func Match(c *gin.Context) {
	appG := app.Gin{C: c}
	validate := validator.New()

	var location Location

	err := appG.C.BindJSON(&location)
	if err != nil {
		appG.Response(http.StatusBadRequest, nil, err)
		return
	}

	err = validate.Struct(location)
	if err != nil {
		appG.Response(http.StatusBadRequest, nil, err)
		return
	}

	request := app.Request{
		Url:         "http://localhost:8000/driver/match",
		ContentType: "application/json",
		Body:        location,
	}

	resp, err := request.Post()
	if err != nil {
		appG.Response(http.StatusBadRequest, nil, err)
		return
	}
	defer resp.Body.Close()

	var data any
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		appG.Response(http.StatusBadRequest, nil, err)
		return
	}
	fmt.Println(data)

	appG.Response(http.StatusOK, data, nil)
}
