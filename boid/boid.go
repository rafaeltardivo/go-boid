package boid

import (
	"math/rand"
	"time"
)

const (
	screenWidth, screenHeight = 640, 360
)

// Boid represents a boid.
type Boid struct {
	Position Vector
	Velocity Vector
	Id       int
}

func (b *Boid) moveOne() {
	newPosition := b.Position.Add(b.Velocity)

	if newPosition.X >= screenWidth || newPosition.X < 0 {
		b.Velocity = Vector{
			X: -b.Velocity.X,
			Y: b.Velocity.Y,
		}
	}

	if newPosition.Y >= screenHeight || newPosition.Y < 0 {
		b.Velocity = Vector{
			X: b.Velocity.X,
			Y: -b.Velocity.Y,
		}
	}

	b.Position = newPosition
}

func (b *Boid) Start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}

}

func CreateBoid(id int) *Boid {
	return &Boid{
		Position: Vector{
			X: rand.Float64() * screenWidth,
			Y: rand.Float64() * screenHeight,
		},
		Velocity: Vector{
			X: (rand.Float64() * 2) - 1.0,
			Y: (rand.Float64() * 2) - 1.0,
		},
		Id: id,
	}
}
