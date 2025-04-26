package dbscan

import (
	"testing"
)

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

func TestGeoPoint_Dimension(t *testing.T) {
	d := GeoPoint{}.Dimension()

	if d != 2 {
		t.Errorf("invalid dimension, expected: 2, got %d\n", d)
	}
}

func TestGeoPoint_AtDimension(t *testing.T) {
	p := GeoPoint{Lng: 1, Lat: 2}

	if p.AtDimension(0) != 1 {
		t.Errorf("invalid first dimension, expected: 1, got %f\n", p.AtDimension(0))
	}

	if p.AtDimension(1) != 2 {
		t.Errorf("invalid second dimension, expected: 2, got %f\n", p.AtDimension(1))
	}

	defer func() {
		if recover() == nil {
			t.Errorf("expected panic on invalid dimension, but did not panic")
		}
	}()

	p.AtDimension(3)
}
