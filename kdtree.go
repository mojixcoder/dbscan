package dbscan

import (
	"math"
	"slices"
)

// Node represents a node in the k-d tree.
type Node struct {
	Point       Point
	Left, Right *Node
}

// KDTree represents the k-d tree structure.
type KDTree struct {
	root *Node
	k    int
}

// newKDTree creates a k-d tree for points of dimension k.
func newKDTree(points []Point) *KDTree {
	k := points[0].Dimension()
	if k <= 0 {
		panic("dimension must be positive")
	}

	root := buildRec(slices.Clone(points), 0, k)
	t := KDTree{k: k, root: root}

	return &t
}

func buildRec(points []Point, depth, k int) *Node {
	if len(points) == 0 {
		return nil
	}

	axis := depth % k
	mid := len(points) / 2

	slices.SortFunc(points, func(a, b Point) int {
		aVal := a.AtDimension(axis)
		bVal := b.AtDimension(axis)

		if aVal == bVal {
			return 0
		} else if aVal < bVal {
			return -1
		}
		return 1
	})

	return &Node{
		Point: points[mid],
		Left:  buildRec(points[:mid], depth+1, k),
		Right: buildRec(points[mid+1:], depth+1, k),
	}
}

// rangeSearch returns all points within the given radius from the query point.
func (t KDTree) rangeSearch(query Point, radius float64) []Point {
	if t.root == nil {
		return []Point{}
	}

	if query.Dimension() != t.k {
		panic("point dimension mismatch")
	}

	var results []Point
	rangeSearchRec(t.root, query, 0, t.k, radius, &results)
	return results
}

// rangeSearchRec recursively collects all points within the given.
func rangeSearchRec(node *Node, query Point, depth, k int, radius float64, results *[]Point) {
	// 1) Exit immediately if there's no node here:
	if node == nil {
		return
	}

	dist := query.DistanceTo(node.Point)
	if dist <= radius {
		*results = append(*results, node.Point)
	}

	d := depth % k
	delta := query.AtDimension(d) - node.Point.AtDimension(d)

	// 2) Always recurse into the “near” child:
	if delta < 0 {
		rangeSearchRec(node.Left, query, depth+1, k, radius, results)
	} else {
		rangeSearchRec(node.Right, query, depth+1, k, radius, results)
	}

	// 3) Only recurse into the “far” child if the hypersphere crosses the split:
	if math.Abs(delta) <= radius {
		if delta < 0 {
			rangeSearchRec(node.Right, query, depth+1, k, radius, results)
		} else {
			rangeSearchRec(node.Left, query, depth+1, k, radius, results)
		}
	}
}
