## Overview
`dbscan` is a lightweight, pure-Go implementation of the DBSCAN clustering algorithm. It works with any type that implements the `Point` interface—which requires only an `ID()` method and a `DistanceTo(Point) float64` method—so you can easily plug in 2D, ND, geospatial, or custom distance metrics.

### ✨ Key features
- k-d Tree Optimization — Accelerates range queries, improving performance on large datasets.
- Supports arbitrary Point types via a simple interface
- Efficient seed-set expansion with enqueue-tracking to avoid duplicates
- Automatic border-point “rescue” of provisional noise
- Returns both per-cluster point lists and per-point cluster labels

### 📦 Installation
```bash
go get github.com/mojixcoder/dbscan
```

### 🚀 Quick Start

``` go
package main

import (
  "fmt"

  "github.com/mojixcoder/dbscan"
)

func main() {
  points := []dbscan.Point{
    dbscan.GeoPoint{PointID: 1, Lat: 1, Lng: 1},
    dbscan.GeoPoint{PointID: 2, Lat: 2, Lng: 2},
    dbscan.GeoPoint{PointID: 3, Lat: 3, Lng: 3},
    dbscan.GeoPoint{PointID: 5, Lat: 5, Lng: 5},
    dbscan.GeoPoint{PointID: 1000, Lat: 1000, Lng: 1000},
    dbscan.GeoPoint{PointID: 1001, Lat: 1001, Lng: 1001},
    dbscan.GeoPoint{PointID: 2000, Lat: 2000, Lng: 2000},
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
```
Example output:
```
### Clusters ###
Cluster 1: [{1 1 1} {2 2 2} {3 3 3} {5 5 5}]
Cluster 2: [{1000 1000 1000} {1001 1001 1001}]
################

### Point Labels ###
ID(1): ClusterID(1)
ID(2): ClusterID(1)
ID(3): ClusterID(1)
ID(5): ClusterID(1)
ID(1000): ClusterID(2)
ID(1001): ClusterID(2)
ID(2000): ClusterID(-1) // Noise
####################
```

### 🧩 API Overview
**Interface:** `Point`  
Your data type must implement:
```go
type Point interface {
	ID() int
	DistanceTo(other Point) float64
	Dimension() int
	AtDimension(int) float64
}
```
✅ Example implementation: [GeoPoint](https://github.com/mojixcoder/dbscan/blob/main/geo_point.go)
