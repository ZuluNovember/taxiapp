package distance

import (
	"math"
)

var rEarth = 6371009.0 //Earth radius in meters

//Gets the distance between two Coordinates.
func Haversine(lon1, lat1, lon2, lat2 float64) int {
	dLon := degreesToRadians(lon1 - lon2)
	dLat := degreesToRadians(lat1 - lat2)

	rlat1 := degreesToRadians(lat1)
	rlat2 := degreesToRadians(lat2)

	angle := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLon/2), 2)*
		math.Cos(rlat1)*math.Cos(rlat2)

	d := 2 * math.Atan2(math.Sqrt(angle), math.Sqrt(1-angle))

	return int(d * rEarth)
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
