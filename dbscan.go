package dbscan

const (
	LabelUndefined = 0
	LabelNoise     = -1
)

type Point interface {
	ID() int
	DistanceTo(other Point) float64
}

// DBScan performs the DBSCAN clustering algorithm on a set of points.
// points    - slice of all points to be clustered.
// eps       - neighborhood radius.
// minPoints - minimum number of points in an eps-neighborhood to form a core point.
func DBScan(points []Point, eps float64, minPoints int) (map[int][]Point, map[int]int) {
	clusterID := 0
	labels := make(map[int]int, len(points))
	queued := make(map[int]bool, len(points))
	clusters := make(map[int][]Point)

	for _, point := range points {
		// Skip if the point is already processed.
		if labels[point.ID()] != LabelUndefined {
			continue
		}

		neighbors := findNeighbors(points, point, eps)

		// Not enough neighbors, mark as noise.
		if len(neighbors) < minPoints {
			labels[point.ID()] = LabelNoise
			continue
		}

		seedSet := make([]Point, 0)
		for _, neighbor := range neighbors {
			if !queued[neighbor.ID()] {
				queued[neighbor.ID()] = true
				seedSet = append(seedSet, neighbor)
			}
		}

		// Found a new cluster, increment the cluster ID.
		clusterID++
		labels[point.ID()] = clusterID
		clusters[clusterID] = []Point{point}

		// Expand the cluster with neighbors.
		for i := 0; i < len(seedSet); i++ {
			neighbor := seedSet[i]

			if labels[neighbor.ID()] == LabelNoise {
				labels[neighbor.ID()] = clusterID
				clusters[clusterID] = append(clusters[clusterID], neighbor)
			}

			// If the neighbor is already labeled (core or border point), skip it.
			if labels[neighbor.ID()] != LabelUndefined {
				continue
			}

			// Mark the neighbor as part of the current cluster.
			labels[neighbor.ID()] = clusterID
			clusters[clusterID] = append(clusters[clusterID], neighbor)

			// If the neighbor is a core point, add its neighbors to the seed set.
			neighborNeighbors := findNeighbors(points, neighbor, eps)
			if len(neighborNeighbors) >= minPoints {
				for _, neighborNeighbor := range neighborNeighbors {
					if queued[neighborNeighbor.ID()] {
						continue
					}
					queued[neighborNeighbor.ID()] = true
					seedSet = append(seedSet, neighborNeighbor)
				}
			}
		}
	}

	return clusters, labels
}

// findNeighbors finds all points in an eps-neighborhood of the query point.
// The query point itself is included.
func findNeighbors(points []Point, query Point, eps float64) []Point {
	neighbors := make([]Point, 0)
	for _, point := range points {
		if query.DistanceTo(point) <= eps {
			neighbors = append(neighbors, point)
		}
	}
	return neighbors
}
