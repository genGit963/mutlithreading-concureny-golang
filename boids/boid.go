package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2d
	velocity Vector2d
	id       int
}

func (b *Boid) calcAccelerationVelocity() Vector2d {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
	avgPosition, avgVelocity := Vector2d{X: 0, Y: 0}, Vector2d{X: 0, Y: 0}
	count := 0.0

	lock.Lock()
	for i := math.Max(lower.X, 0); i < math.Min(upper.X, screenWidth); i++ {
		for j := math.Max(lower.Y, 0); j < math.Min(upper.Y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
					count = count + 1
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].velocity)
				}
			}
		}
	}
	lock.Unlock()

	accelarationVelocity := Vector2d{X: 0, Y: 0}
	if count > 0 {
		avgPosition, avgVelocity = avgPosition.DivisionV(count), avgVelocity.DivisionV(count)
		accelarationAlignment := avgVelocity.Subtract(b.velocity).MultiplyV(adjRate)
		accelarationCohesion := avgPosition.Subtract(b.position).MultiplyV(adjRate)

		accelarationVelocity = accelarationVelocity.Add(accelarationAlignment).Add(accelarationCohesion)
	}
	return accelarationVelocity

}

func (b *Boid) moveOne() {
	accelation := b.calcAccelerationVelocity()
	lock.Lock()
	b.velocity = b.velocity.Add(accelation).limit(-1, 1)
	boidMap[int(b.position.X)][int(b.position.Y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.X)][int(b.position.Y)] = b.id

	next := b.position.Add(b.velocity)
	if next.X >= screenWidth || next.X < 0 {
		b.velocity = Vector2d{X: -b.velocity.X, Y: b.velocity.Y}
	}

	if next.Y >= screenHeight || next.Y < 0 {
		b.velocity = Vector2d{X: b.velocity.X, Y: -b.velocity.Y}
	}
	lock.Unlock()
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
