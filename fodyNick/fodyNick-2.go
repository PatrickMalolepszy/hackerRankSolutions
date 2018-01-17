package fodyNick

//import "math"

// same box class
// write a function given a
// set of points to calculate the minimum bounding
// box for the set of points

type box struct {
	minX, minY, maxX, maxY float32
}
//
//// set of points given as slice of pairs
//func findABox(points [][]int32) *box {
//	if len(points) == 0 {
//		return &box{}
//	}
//	box := &box{
//		minX: math.MaxInt32,
//		minY: math.MaxInt32,
//		maxX: math.MinInt32,
//		maxY: math.MinInt32,
//	}
//
//	for i := range points {
//		if points[i][0] < box.minX {
//			box.minX = points[i][0]
//		}
//		if points[i][1] < box.minY {
//			box.minY = points[i][1]
//		}
//		if points[i][0] > box.maxX {
//			box.maxX = points[i][0]
//		}
//		if points[i][1] > box.maxY {
//			box.maxY = points[i][1]
//		}
//	}
//
//	return box
//}