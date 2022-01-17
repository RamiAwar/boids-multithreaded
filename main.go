package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const screen_width, screen_height = 1440, 900
const boid_count = 10000

var GREEN = color.RGBA{10, 255, 50, 255}

var boids [boid_count]*Boid

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), GREEN)
		screen.Set(int(boid.position.x-1), int(boid.position.y), GREEN)
		screen.Set(int(boid.position.x), int(boid.position.y-1), GREEN)
		screen.Set(int(boid.position.x), int(boid.position.y+1), GREEN)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screen_width, screen_height
}

func main() {
	for i := 0; i < boid_count; i++ {
		boid := CreateBoid(i)
		boids[i] = &boid
		go boid.Start()
	}
	ebiten.SetWindowSize(screen_width, screen_height)
	ebiten.SetWindowTitle("Boids Multithreaded")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
