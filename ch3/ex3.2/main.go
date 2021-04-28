package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            //画布大小
	cells         = 100                 //单元格的个数
	xyrange       = 30.0                //坐标轴的范围(-xyrnage..+xyrange)
	xyscale       = width / 2 / xyrange //x或y轴上每个单位长度的像素
	zscale        = height * 0.4        //z轴上每个单位长度的像素
	angle         = math.Pi / 6         //x、y轴的角度(=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°),cos(30°)

func main() {
	file, err := os.Create("test.svg")
	if err != nil {
		fmt.Errorf("err:%v", err)
		return
	}
	zMin, zMax := minMax()
	fmt.Fprintf(file, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			} else {
				fmt.Fprintf(file, "<polygon style='stroke: %s;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					color(i, j, zMin, zMax), ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintf(file, "</svg>")
}

// minmax返回给定x和y的最小值/最大值并假设为方域的z的最小值和最大值。
func minMax() (min, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xoff := 0; xoff <= 1; xoff++ {
				for yoff := 0; yoff <= 1; yoff++ {
					x := xyrange * (float64(i+xoff)/cells - 0.5)
					y := xyrange * (float64(j+yoff)/cells - 0.5)
					z := f(x, y)
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}
		}
	}
	return min, max
}

func color(i, j int, zmin, zmax float64) string {
	min := math.NaN()
	max := math.NaN()
	for xoff := 0; xoff <= 1; xoff++ {
		for yoff := 0; yoff <= 1; yoff++ {
			x := xyrange * (float64(i+xoff)/cells - 0.5)
			y := xyrange * (float64(j+yoff)/cells - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	color := ""
	if math.Abs(max) > math.Abs(min) {
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}
		color = fmt.Sprintf("#%02x0000", int(red))
	} else {
		blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
		if blue > 255 {
			blue = 255
		}
		color = fmt.Sprintf("#0000%02x", int(blue))
	}
	return color
}

func corner(i, j int) (float64, float64) {
	//求出网格单元(i,j)的顶点坐标(x,y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//计算曲面高度z
	//z := f(x, y)
	//z := eggbox(x, y)
	z := saddle(x, y)
	//将(x, y, z)等角投射到二维SVG绘图平面上,坐标是(sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 { //雪堆
	r := math.Hypot(x, y) //到(0,0)的距离
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 { //鸡蛋盒
	r := 0.2 * (math.Cos(x) + math.Cos(y))
	return r
}

func saddle(x, y float64) float64 { //马鞍
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	r := y*y/a2 - x*x/b2
	return r
}
