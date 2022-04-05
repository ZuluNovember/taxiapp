package api

import (
	"net/http"

	"github.com/ZuluNovember/taxiapp/driver/models"
	"github.com/ZuluNovember/taxiapp/driver/pkg/app"
	"github.com/ZuluNovember/taxiapp/driver/service/csv"
	"github.com/ZuluNovember/taxiapp/driver/service/matching"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Location struct {
	Lon float64 `json:"lon" validate:"gte=-180.0,lte=180.0,required"`
	Lat float64 `json:"lat" validate:"gte=-90.0,lte=90.0,required"`
}

// @Summary Import coordinates in Coordinates.csv file to mongo
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /driver/import [get]
func ImportCsv(c *gin.Context) {
	appG := app.Gin{C: c}

	err := csv.CsvCoordinatesToMongo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, nil, err)
		return
	}
	appG.Response(http.StatusOK, nil, nil)
}

// @Summary Add a driver to mongodb with lon and lat values
// @Accept json
// @Produce  json
// @Param coordinate body api.Location true "Coordinate"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /driver/add [post]
func AddDriver(c *gin.Context) {
	appG := app.Gin{C: c}
	var location Location

	validate := validator.New()

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

	res, err := models.AddDriver(location.Lon, location.Lat)
	if err != nil {
		appG.Response(http.StatusBadRequest, nil, err)
		return
	}

	appG.Response(http.StatusOK, res, nil)
}

// TODO: Function doesn't return expected result in swagger but works in Postman.
// @Summary Find a driver close to given coordinates
// @Tags driver
// @Accept json
// @Produce  json
// @Param coordinate body api.Location true "Coordinate"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /driver/match [post]
func FindDriver(c *gin.Context) {
	appG := app.Gin{C: c}
	var location Location

	validate := validator.New()

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

	rider := matching.Rider{
		Lon: location.Lon,
		Lat: location.Lat,
	}
	match, err := rider.Match()
	if err != nil {
		appG.Response(http.StatusBadRequest, nil, err)
		return
	}

	appG.Response(http.StatusOK, match, nil)
}
