package dbscan

import "testing"

func TestGeoPoint_ID(t *testing.T) {
	p := GeoPoint{PointID: 1}

	if p.ID() != 1 {
		t.Errorf("invalid ID, expected: %d, actual %d\n", 1, p.ID())
	}
}

func TestGeoPoint_DistanceTo(t *testing.T) {
	p1 := GeoPoint{Lat: 1, Lng: 1}
	p2 := GeoPoint{Lat: 1, Lng: 3}

	distance := p1.DistanceTo(p2)

	if distance != 2 {
		t.Errorf("invalid distance, expected: %f, actual %f\n", 2.0, distance)
	}
}
