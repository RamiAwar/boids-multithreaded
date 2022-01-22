package main

import (
	"image"
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const screen_width, screen_height = 640, 360
const boid_count = 1300
const view_radius = 13
const boid_size = 4
const acceleration_rate = 0.03
const cohesion_rate = 0.015
const separation_rate = 0.03

const cell_size = 40

var spatialHashGrid = NewSpatialHash(cell_size)

var emptyImage = ebiten.NewImage(3, 3)
var emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

var GREEN = color.RGBA{10, 255, 50, 255}

var boids [boid_count]*Boid

var lock = sync.RWMutex{}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		drawBoid(screen, boid)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screen_width, screen_height
}

func drawBoid(screen *ebiten.Image, boid *Boid) {
	var path vector.Path

	heading := boid.velocity.Normalize()

	heading = heading.MulV(boid_size / 3)
	path.MoveTo(float32(boid.position.x-heading.y), float32(boid.position.y+heading.x)) // sides
	path.LineTo(float32(boid.position.x+heading.y), float32(boid.position.y-heading.x)) // sides
	heading = heading.MulV(boid_size)
	path.LineTo(float32(boid.position.x+heading.x), float32(boid.position.y+heading.y)) // tip

	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}
	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0xdb / float32(0xff)
		vs[i].ColorG = 0x56 / float32(0xff)
		vs[i].ColorB = 0x20 / float32(0xff)
	}

	screen.DrawTriangles(vs, is, emptySubImage, op)

	// screen.Set(int(boid.position.x+1), int(boid.position.y), GREEN)
	// screen.Set(int(boid.position.x-1), int(boid.position.y), GREEN)
	// screen.Set(int(boid.position.x), int(boid.position.y-1), GREEN)
	// screen.Set(int(boid.position.x), int(boid.position.y+1), GREEN)
}

func main() {
	emptyImage.Fill(color.White)

	lock.Lock()
	for i := 0; i < boid_count; i++ {
		boid := CreateBoid(i)
		boids[i] = &boid
		spatialHashGrid.Add(boid)
		go boid.Start()
	}
	lock.Unlock()

	ebiten.SetWindowSize(screen_width*2, screen_height*2)
	ebiten.SetWindowTitle("Boids Multithreaded")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
