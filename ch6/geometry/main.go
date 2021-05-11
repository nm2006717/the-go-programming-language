package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

// traditional function

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//	same thing,but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	var p, q = Point{
		X: 1,
		Y: 1,
	}, Point{
		X: 4,
		Y: 5,
	}
	dis := Distance(p, q)
	fmt.Printf("%v to %v distance %f\n", p, q, dis)

	dis2 := p.Distance(q)
	fmt.Printf("%v to %v distance %f\n", p, q, dis2)

	//	计算三角形的周长
	var perim = Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	r := &Point{
		X: 1,
		Y: 2,
	}
	r.ScaleBy(2)
	fmt.Println(*r)

	p2 := Point{1, 2}
	pptr := &p2
	pptr.ScaleBy(2)
	fmt.Println(p2) // "{2, 4}"

	p3 := Point{1, 2}
	(&p3).ScaleBy(2)
	fmt.Println(p3) // "{2, 4}"

	p4 := Point{1, 2}
	p4.ScaleBy(2)
	fmt.Println(p4) // "{2, 4}"

}
