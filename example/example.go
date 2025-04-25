package main

import (
	"fmt"

	"github.com/mojixcoder/dbscan"
)

func main() {
	points := []dbscan.Point{
		dbscan.GeoPoint{
			PointID: 1,
			Lat:     1,
			Lng:     1,
		},
		dbscan.GeoPoint{
			PointID: 2,
			Lat:     2,
			Lng:     2,
		},
		dbscan.GeoPoint{
			PointID: 3,
			Lat:     3,
			Lng:     3,
		},
		dbscan.GeoPoint{
			PointID: 5,
			Lat:     5,
			Lng:     5,
		},
		dbscan.GeoPoint{
			PointID: 1000,
			Lat:     1000,
			Lng:     1000,
		},
		dbscan.GeoPoint{
			PointID: 1001,
			Lat:     1001,
			Lng:     1001,
		},
		dbscan.GeoPoint{
			PointID: 2000,
			Lat:     2000,
			Lng:     2000,
		},
	}

	eps := 5.0
	minPoints := 2

	result := dbscan.DBScan(points, eps, minPoints)

	fmt.Println("### Clusters ###")
	for clusterID, cluster := range result.Clusters {
		fmt.Printf("Cluster %d: %v\n", clusterID, cluster)
	}
	fmt.Println("################")
	fmt.Println("")

	fmt.Println("### Point Labels ###")
	for pointID, label := range result.Labels {
		fmt.Printf("ID(%d): ClusterID(%d)\n", pointID, label)
	}
	fmt.Println("####################")
}
