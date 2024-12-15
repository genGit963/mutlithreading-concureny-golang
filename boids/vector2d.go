package main

import "math"

type Vector2d struct {
	X float64
	Y float64
}

func (v1 Vector2d) Add(v2 Vector2d) Vector2d {
	return Vector2d{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func (v1 Vector2d) Sub(v2 Vector2d) Vector2d {
	return Vector2d{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

func (v1 Vector2d) Multiply(v2 Vector2d) Vector2d {
	return Vector2d{X: v1.X * v2.X, Y: v1.Y * v2.Y}
}

func (v1 Vector2d) AddV(d float64) Vector2d {
	return Vector2d{X: v1.X + d, Y: v1.Y + d}
}

func (v1 Vector2d) SubV(d float64) Vector2d {
	return Vector2d{X: v1.X - d, Y: v1.Y - d}
}

func (v1 Vector2d) MultiplyV(d float64) Vector2d {
	return Vector2d{X: v1.X * d, Y: v1.Y * d}
}

func (v1 Vector2d) DivisionV(d float64) Vector2d {
	return Vector2d{X: v1.X / d, Y: v1.Y / d}
}

func (v1 Vector2d) limit(lower, upper float64) Vector2d {
	return Vector2d{
		X: math.Min(math.Max(v1.X, lower), upper),
		Y: math.Min(math.Max(v1.Y, lower), upper),
	}
}

func (v1 Vector2d) Distance(v2 Vector2d) float64 {
	return math.Sqrt(math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2))
}
