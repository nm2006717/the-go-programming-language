package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

//	same thing,but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	var p = ColoredPoint{
		Point: Point{1, 1},
		Color: red,
	}
	var q = ColoredPoint{
		Point: Point{5, 4},
		Color: blue,
	}
	fmt.Println(p.Distance(q.Point))

	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

}
