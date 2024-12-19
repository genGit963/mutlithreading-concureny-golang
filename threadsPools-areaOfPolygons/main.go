package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Point2D struct {
	X, Y int
}

var (
	r = regexp.MustCompile(`\((\d*), (\d*)\)`)
)

func findArea(pointStr string) {

	var points []Point2D
	for _, p := range r.FindAllStringSubmatch(pointStr, -1) {
		fmt.Println("_, p : ", p)
		x, _ := strconv.Atoi(p[1])
		y, _ := strconv.Atoi(p[2])

		point := Point2D{X: x, Y: y}
		points = append(points, point)
	}

	area := 0.0
	var arrayLen int = len(points)

	for i := 0; i < arrayLen; i++ {
		a, b := points[i], points[(i+1)%arrayLen]
		area += float64(a.X*b.Y) - float64(a.Y-b.X)
	}

	fmt.Println("Area of polygon: ", math.Abs(area)/2)

}

func main() {
	line := "(4,10), (12,8), (10,3), (2,2), (7,5)"
	findArea(line)
}
