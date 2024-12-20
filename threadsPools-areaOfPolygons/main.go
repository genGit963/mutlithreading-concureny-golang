package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Point2D struct {
	X, Y int
}

const (
// totalWorkerThreads = 10 // fixed no of threads
)

var (
	r = regexp.MustCompile(`\((\d*), (\d*)\)`)
	// waitGroup          = sync.WaitGroup{}
	totalWorkerThreads = runtime.NumCPU() // dynamic number of own sys threads
)

func findArea(pointStr string) float64 {
	// func findArea(inputChannel chan string) {

	// for pointStr := range inputChannel {
	var points []Point2D
	for _, p := range r.FindAllStringSubmatch(pointStr, -1) {
		// fmt.Println("_, p : ", p)
		x, _ := strconv.Atoi(p[1])
		y, _ := strconv.Atoi(p[2])

		point := Point2D{X: x, Y: y}
		points = append(points, point)
	}

	area := 0.0
	var arrayLen int = len(points)

	for i := 0; i < arrayLen; i++ {
		a, b := points[i], points[(i+1)%arrayLen]
		area += float64(a.X*b.Y) - float64(a.Y*b.X)
	}

	// fmt.Println("Area of polygon: ", math.Abs(area)/2)
	// }
	// waitGroup.Done()
	return math.Abs(area / 2)
}

func main() {

	// read file
	data, _ := os.ReadFile("/Users/maheshbogati/Desktop/multi-threading-golang/threadsPools-areaOfPolygons/input-text.txt")
	text := string(data)

	// inputChan := make(chan string, 50) // buffer-size: 50
	// // worker threads
	// for i := 0; i < totalWorkerThreads; i++ {
	// 	go findArea(inputChan)
	// }
	// waitGroup.Add(totalWorkerThreads)

	startTime := time.Now()
	var results []float64
	for _, line := range strings.Split(text, "\n") {
		// inputChan <- line
		// fmt.Println("passed to buffer: ", index)
		results = append(results, findArea(line))
	}
	// close(inputChan)

	// wait
	// waitGroup.Wait()

	elapsed := time.Since(startTime)
	fmt.Println("Done in: ", elapsed)
	fmt.Println("no of line executed : ", len(results))
	// fmt.Println("Total Active Threads: ", totalWorkerThreads)
}
