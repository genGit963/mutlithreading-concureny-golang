package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2d
	velocity Vector2d
	id       int
}

func (b *Boid) moveOne() {
	b.position = b.position.Add(b.velocity)
	next := b.position.Add(b.velocity)

	if next.X >= screenWidth || next.X < 0 {
		b.velocity = Vector2d{X: -b.velocity.X, Y: b.velocity.Y}
	}

	if next.Y >= screenHeight || next.Y < 0 {
		b.velocity = Vector2d{X: b.velocity.X, Y: -b.velocity.Y}
	}
	// fmt.Println("velocity: ", b.id, b.velocity)
}

func (b *Boid) start() {

	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}

}

func createBoid(bid int) {

	b := Boid{
		position: Vector2d{X: rand.Float64() * screenWidth, Y: rand.Float64() * screenHeight},
		velocity: Vector2d{X: (rand.Float64() * 2) - 1.0, Y: (rand.Float64() * 2) - 1.0},
		id:       bid,
	}
	boids[bid] = &b
	boidMap[int(b.position.X)][int(b.position.Y)] = b.id

	go b.start()
}
