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
	b.Position = b.Position.Add(b.Velocity)
	next := b.Position.Add(b.Velocity)

	if next.X >= screenWidth || next.X < 0 {
		b.Velocity = Vector{
			X: -b.Velocity.X,
			Y: b.Velocity.Y,
		}
	}

	if next.Y >= screenHeight || next.Y < 0 {
		b.Velocity = Vector{
			X: b.Velocity.X,
			Y: -b.Velocity.Y,
		}
	}
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
