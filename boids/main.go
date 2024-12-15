package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 440, 240
	boidCount                 = 10000
	viewRadius                = 13
	adjRate                   = 0.015
)

var (
	green   = color.RGBA{10, 255, 10, 255}
	boids   [boidCount]*Boid
	boidMap [screenWidth + 2][screenHeight + 2]int
	lock    = sync.Mutex{}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.X+1), int(boid.position.Y), green)
		screen.Set(int(boid.position.X-1), int(boid.position.Y), green)
		screen.Set(int(boid.position.X+1), int(boid.position.Y-1), green)
		screen.Set(int(boid.position.X+1), int(boid.position.Y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {

	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
