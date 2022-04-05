package distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHaversine(t *testing.T) {
	testCases := []struct {
		lon1, lat1, lon2, lat2 float64
		want                   int
	}{
		{14.1234, -14.14141, 13.9494, 14.2121, 3152830},
		{-123.241341, 83.2234234, -123.12124, 83.21342, 1930},
	}
	for _, tC := range testCases {
		got := Haversine(tC.lon1, tC.lat1, tC.lon2, tC.lat2)
		assert.InDelta(t, tC.want, got, 1.0)
	}
}
