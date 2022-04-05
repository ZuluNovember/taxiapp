package matching

import (
	"errors"

	"github.com/ZuluNovember/taxiapp/driver/models"
	"github.com/ZuluNovember/taxiapp/driver/pkg/distance"
)

type Rider struct {
	Lon float64
	Lat float64
}

type MatchResult struct {
	Driver   models.Driver
	Distance int
}

func (r Rider) Match() (*MatchResult, error) {
	var res MatchResult

	drivers, err := models.FindDrivers(r.Lon, r.Lat)
	if err != nil {
		return nil, err
	}
	if len(drivers) == 0 {
		return nil, errors.New("No drivers nearby")
	}

	var d models.Driver = drivers[0]
	dist := distance.Haversine(d.Location.Coordinates[0], d.Location.Coordinates[1], r.Lon, r.Lat)

	res.Distance = dist
	res.Driver = d

	return &res, nil
}
