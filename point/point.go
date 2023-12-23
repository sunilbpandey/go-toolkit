package point

import "github.com/sunilbpandey/go-toolkit/intutils"

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p Point) Distance(q Point) int {
	return intutils.Abs(p.X-q.X) + intutils.Abs(p.Y-q.Y)
}
