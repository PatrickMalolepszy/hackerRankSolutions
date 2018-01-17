package fodyNick
// first, you have a box with 2 points, each point contains x and y
// the box points are labeled min and max
// the x and y in min are less than the x and y in max respectively
// write a function that given a new point,
// you decide whether it is in or out of the box

func outOfTheBox(minX, minY, maxX, maxY, x, y int) bool {
	if minX <= x && x <= maxX && minY <= y && y <= maxY {
		return true
	} else {
		return false
	}
}