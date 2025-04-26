package dbscan

import (
	"math"
	"slices"
	"testing"
)

type testPoint struct {
	GeoPoint
}

func (testPoint) Dimension() int {
	return -1
}

func TestKDTree_rangeSearch(t *testing.T) {
	pts := []Point{
		GeoPoint{Lat: 0, Lng: 0},
		GeoPoint{Lat: 0, Lng: 0},
		GeoPoint{Lat: 2, Lng: 2},
		GeoPoint{Lat: 1, Lng: 1},
		GeoPoint{Lat: 3, Lng: 3},
		GeoPoint{Lat: 4, Lng: 4},
	}

	testCases := []struct {
		name     string
		query    Point
		radius   float64
		expected []Point
		tree     *KDTree
	}{
		{
			name:   "search_with_neighbors",
			query:  GeoPoint{Lat: 1, Lng: 1},
			radius: math.Sqrt(2),
			expected: []Point{
				GeoPoint{Lat: 2, Lng: 2},
				GeoPoint{Lat: 1, Lng: 1},
				GeoPoint{Lat: 0, Lng: 0},
				GeoPoint{Lat: 0, Lng: 0},
			},
			tree: newKDTree(pts),
		},
		{
			name:   "no_neighbors",
			query:  GeoPoint{Lat: 1, Lng: 1},
			radius: 0,
			tree:   newKDTree(pts),
			expected: []Point{
				GeoPoint{Lat: 1, Lng: 1},
			},
		},
		{
			name:   "all_points",
			query:  GeoPoint{Lat: 1, Lng: 1},
			radius: math.Sqrt(18),
			tree:   newKDTree(pts),
			expected: []Point{
				GeoPoint{Lat: 0, Lng: 0},
				GeoPoint{Lat: 2, Lng: 2},
				GeoPoint{Lat: 1, Lng: 1},
				GeoPoint{Lat: 0, Lng: 0},
				GeoPoint{Lat: 4, Lng: 4},
				GeoPoint{Lat: 3, Lng: 3},
			},
		},
		{
			name:     "nil_root",
			query:    GeoPoint{Lat: 1, Lng: 1},
			radius:   100,
			expected: []Point{},
			tree:     &KDTree{root: nil},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.tree.rangeSearch(tc.query, tc.radius)

			slices.SortFunc(result, func(a, b Point) int {
				lngA := a.AtDimension(0)
				lngB := b.AtDimension(0)

				return int(lngA - lngB)
			})

			slices.SortFunc(tc.expected, func(a, b Point) int {
				lngA := a.AtDimension(0)
				lngB := b.AtDimension(0)

				return int(lngA - lngB)
			})

			if !slices.Equal(result, tc.expected) {
				t.Errorf("wrong range search, expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestNewKDTree(t *testing.T) {
	pts := []Point{GeoPoint{}}
	tree := newKDTree(pts)

	if tree.root == nil {
		t.Error("expected non-nil root for single point")
	}

	if (tree.root.Point.(GeoPoint) != GeoPoint{}) {
		t.Error("expected non-nil root for single point")
	}
}

func TestNewKDTree_Panics_Invalid_Dimension(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("expected panic on invalid dimension, but did not panic")
		}
	}()

	newKDTree([]Point{testPoint{}})
}

func TestBuildRec_Nil_Check(t *testing.T) {
	node := buildRec(nil, 0, 1)

	if node != nil {
		t.Error("expected nil node for single point")
	}
}

func TestKDTree_rangeSearch_Invalid_Dimension(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("expected panic on invalid dimension, but did not panic")
		}
	}()

	tree := KDTree{k: 1, root: &Node{}}
	tree.rangeSearch(GeoPoint{}, 5)
}
