package dbscan

import (
	"fmt"
	"math"
)

// GeoPoint represents a point in geographical coordinates.
// It implements the Point interface for use in DBSCAN clustering.
// DistanceTo method is implemented using Euclidean distance algorithm.
type GeoPoint struct {
	PointID int
	Lat     float64
	Lng     float64
}

// ID returns point's ID.
func (p GeoPoint) ID() int {
	return p.PointID
}

// DistanceTo returns the Euclidean distance between two points.
func (p GeoPoint) DistanceTo(other Point) float64 {
	otherGeoPoint := other.(GeoPoint)
	deltaLat := p.Lat - otherGeoPoint.Lat
	deltaLng := p.Lng - otherGeoPoint.Lng
	return math.Sqrt((deltaLat * deltaLat) + (deltaLng * deltaLng))
}

func (p GeoPoint) Dimension() int {
	return 2
}

func (p GeoPoint) AtDimension(d int) float64 {
	switch d {
	case 0:
		return p.Lng
	case 1:
		return p.Lat
	default:
		panic("invalid dimension: " + fmt.Sprint(d))
	}
}
