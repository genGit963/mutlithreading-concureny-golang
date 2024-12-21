package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Point2D struct {
	X, Y int
}

var (
	r                  = regexp.MustCompile(`\\((\\d*), (\\d*)\\)`)
	waitGroup          = sync.WaitGroup{}
	totalWorkerThreads = runtime.NumCPU() // dynamic number of threads based on system CPU
)

// Worker function to compute the area of polygons
func findArea(inputChannel chan string, resultsChannel chan float64) {
	defer waitGroup.Done() // Ensure WaitGroup counter is decremented when worker finishes

	for pointStr := range inputChannel {
		var points []Point2D
		for _, p := range r.FindAllStringSubmatch(pointStr, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])

			points = append(points, Point2D{X: x, Y: y})
		}

		area := 0.0
		arrayLen := len(points)

		for i := 0; i < arrayLen; i++ {
			a, b := points[i], points[(i+1)%arrayLen]
			area += float64(a.X*b.Y) - float64(a.Y*b.X)
		}

		// Send the computed area to the results channel
		resultsChannel <- math.Abs(area / 2)
	}
}

func main() {
	// Read input file
	data, _ := os.ReadFile("/Users/maheshbogati/Desktop/multi-threading-golang/threadsPools-areaOfPolygons/input-text.txt")
	text := string(data)

	inputChan := make(chan string, 50)   // Buffer-size for input
	resultChan := make(chan float64, 50) // Buffer-size for results

	// Start worker threads
	for i := 0; i < totalWorkerThreads; i++ {
		waitGroup.Add(1)
		go findArea(inputChan, resultChan)
	}

	// Start timing
	startTime := time.Now()

	// Send data to input channel
	go func() {
		for _, line := range strings.Split(text, "\n") {
			inputChan <- line
		}
		close(inputChan) // Close input channel after sending all lines
	}()

	// Close the results channel after all workers finish
	go func() {
		waitGroup.Wait() // Wait for all workers to complete
		close(resultChan)
	}()

	// Collect results
	var results []float64
	for result := range resultChan {
		results = append(results, result)
	}

	// Measure elapsed time
	elapsed := time.Since(startTime)
	fmt.Println("Done in:", elapsed)
	fmt.Println("Number of polygons processed:", len(results))
	fmt.Println("Total Active Threads:", totalWorkerThreads)
}
