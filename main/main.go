package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rafaeltardivo/go-boid/boid"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
)

var (
	green = color.RGBA{R: 10, G: 255, B: 50, A: 255}
	boids [boidCount]*boid.Boid
)

// Game represents the game.
type Game struct{}

// Update updates the game.
func (g *Game) Update() error {
	return nil
}

// Draw draws the game.
func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.Position.X+1), int(boid.Position.Y), green)
		screen.Set(int(boid.Position.X-1), int(boid.Position.Y), green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y-1), green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {

	for i := 0; i < boidCount; i++ {
		boids[i] = boid.CreateBoid(i)
		go boids[i].Start()
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal((err))
	}
}
