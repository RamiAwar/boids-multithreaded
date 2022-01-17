package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func CreateBoid(id int) Boid {
	return Boid{
		position: Vector2D{x: rand.Float64() * screen_width, y: rand.Float64() * screen_height},
		velocity: Vector2D{x: rand.Float64()*2 - 1.0, y: rand.Float64()*2 - 1},
		id:       id,
	}
}

func (b *Boid) Start() {
	for {
		b.MoveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) MoveOne() {
	b.position.Add(b.velocity)
	if b.position.x >= screen_width || b.position.x < 0 {
		b.velocity.Mul(Vector2D{-1.0, 1.0})
	}

	if b.position.y >= screen_height || b.position.y < 0 {
		b.velocity.Mul(Vector2D{1.0, -1.0})
	}
}
