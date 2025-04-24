package dbscan

import (
	"maps"
	"slices"
	"testing"
)

func TestDBScan(t *testing.T) {
	points := []Point{
		GeoPoint{
			PointID: 1,
			Lat:     1,
			Lng:     1,
		},
		GeoPoint{
			PointID: 2,
			Lat:     2,
			Lng:     2,
		},
		GeoPoint{
			PointID: 3,
			Lat:     3,
			Lng:     3,
		},
		GeoPoint{
			PointID: 5,
			Lat:     5,
			Lng:     5,
		},
		GeoPoint{
			PointID: 1000,
			Lat:     1000,
			Lng:     1000,
		},
		GeoPoint{
			PointID: 1001,
			Lat:     1001,
			Lng:     1001,
		},
		GeoPoint{
			PointID: 2000,
			Lat:     2000,
			Lng:     2000,
		},
	}

	testCases := []struct {
		name             string
		minPoints        int
		eps              float64
		expectedClusters map[int][]Point
		expectedLabels   map[int]int
	}{
		{
			name:      "one_cluster",
			minPoints: 3,
			eps:       5,
			expectedClusters: map[int][]Point{
				1: {
					GeoPoint{
						PointID: 1,
						Lat:     1,
						Lng:     1,
					},
					GeoPoint{
						PointID: 2,
						Lat:     2,
						Lng:     2,
					},
					GeoPoint{
						PointID: 3,
						Lat:     3,
						Lng:     3,
					},
					GeoPoint{
						PointID: 5,
						Lat:     5,
						Lng:     5,
					},
				},
			},
			expectedLabels: map[int]int{
				1:    1,
				2:    1,
				3:    1,
				5:    1,
				1000: -1,
				1001: -1,
				2000: -1,
			},
		},
		{
			name:      "one_cluster_noise_first",
			minPoints: 4,
			eps:       3,
			expectedLabels: map[int]int{
				1:    1,
				2:    1,
				3:    1,
				5:    1,
				1000: -1,
				1001: -1,
				2000: -1,
			},
			expectedClusters: map[int][]Point{
				1: {
					GeoPoint{
						PointID: 3,
						Lat:     3,
						Lng:     3,
					},
					GeoPoint{
						PointID: 1,
						Lat:     1,
						Lng:     1,
					},
					GeoPoint{
						PointID: 2,
						Lat:     2,
						Lng:     2,
					},
					GeoPoint{
						PointID: 5,
						Lat:     5,
						Lng:     5,
					},
				},
			},
		},
		{
			name:      "two_cluster",
			minPoints: 2,
			eps:       5,
			expectedLabels: map[int]int{
				1:    1,
				2:    1,
				3:    1,
				5:    1,
				1000: 2,
				1001: 2,
				2000: -1,
			},
			expectedClusters: map[int][]Point{
				1: {
					GeoPoint{
						PointID: 1,
						Lat:     1,
						Lng:     1,
					},
					GeoPoint{
						PointID: 2,
						Lat:     2,
						Lng:     2,
					},
					GeoPoint{
						PointID: 3,
						Lat:     3,
						Lng:     3,
					},
					GeoPoint{
						PointID: 5,
						Lat:     5,
						Lng:     5,
					},
				},
				2: {
					GeoPoint{
						PointID: 1000,
						Lat:     1000,
						Lng:     1000,
					},
					GeoPoint{
						PointID: 1001,
						Lat:     1001,
						Lng:     1001,
					},
				},
			},
		},
		{
			name:      "all_cluster",
			minPoints: 1,
			eps:       5,
			expectedLabels: map[int]int{
				1:    1,
				2:    1,
				3:    1,
				5:    1,
				1000: 2,
				1001: 2,
				2000: 3,
			},
			expectedClusters: map[int][]Point{
				1: {
					GeoPoint{
						PointID: 1,
						Lat:     1,
						Lng:     1,
					},
					GeoPoint{
						PointID: 2,
						Lat:     2,
						Lng:     2,
					},
					GeoPoint{
						PointID: 3,
						Lat:     3,
						Lng:     3,
					},
					GeoPoint{
						PointID: 5,
						Lat:     5,
						Lng:     5,
					},
				},
				2: {
					GeoPoint{
						PointID: 1000,
						Lat:     1000,
						Lng:     1000,
					},
					GeoPoint{
						PointID: 1001,
						Lat:     1001,
						Lng:     1001,
					},
				},
				3: {
					GeoPoint{
						PointID: 2000,
						Lat:     2000,
						Lng:     2000,
					},
				},
			},
		},
		{
			name:      "all_noise_because_of_min_points",
			minPoints: 5,
			eps:       5,
			expectedLabels: map[int]int{
				1:    -1,
				2:    -1,
				3:    -1,
				5:    -1,
				1000: -1,
				1001: -1,
				2000: -1,
			},
			expectedClusters: map[int][]Point{},
		},
		{
			name:      "all_noise_because_of_distance",
			minPoints: 2,
			eps:       0,
			expectedLabels: map[int]int{
				1:    -1,
				2:    -1,
				3:    -1,
				5:    -1,
				1000: -1,
				1001: -1,
				2000: -1,
			},
			expectedClusters: map[int][]Point{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clusters, labels := DBScan(points, tc.eps, tc.minPoints)

			if !maps.Equal(tc.expectedLabels, labels) {
				t.Errorf("labels are not equal, expected: %v, actual: %v\n", tc.expectedLabels, labels)
			}

			if !maps.EqualFunc(tc.expectedClusters, clusters, func(v1, v2 []Point) bool {
				return slices.Equal(v1, v2)
			}) {
				t.Errorf("clusters are not equal, expected: %v, actual: %v\n", tc.expectedClusters, clusters)
			}
		})
	}
}

func TestFindNeighbors(t *testing.T) {
	points := []Point{
		GeoPoint{
			PointID: 1,
			Lat:     1,
			Lng:     1,
		},
		GeoPoint{
			PointID: 2,
			Lat:     2,
			Lng:     2,
		},
		GeoPoint{
			PointID: 3,
			Lat:     3,
			Lng:     3,
		},
		GeoPoint{
			PointID: 1000,
			Lat:     1000,
			Lng:     1000,
		},
	}

	expectedNeighbors := points[:3]
	neighbors := findNeighbors(points, points[1], 5)
	if !slices.Equal(expectedNeighbors, neighbors) {
		t.Errorf("neighbors are not equal, expected: %v, actual: %v\n", expectedNeighbors, neighbors)
	}

	expectedNeighbors = points[3:]
	neighbors = findNeighbors(points, points[3], 5)
	if !slices.Equal(expectedNeighbors, neighbors) {
		t.Errorf("neighbors are not equal, expected: %v, actual: %v\n", expectedNeighbors, neighbors)
	}
}
