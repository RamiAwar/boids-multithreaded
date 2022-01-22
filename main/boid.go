package main

import (
	"math"
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
		velocity: Vector2D{x: rand.Float64() - 0.5, y: rand.Float64() - 0.5},
		id:       id,
	}
}

func (b *Boid) Start() {
	for {
		b.MoveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) Avoid(width, height float64) Vector2D {
	accel := Vector2D{0.0, 0.0}
	if b.position.x <= view_radius {
		accel.x = 1 / b.position.x
	} else if b.position.x >= width-view_radius {
		accel.x = -1 / math.Abs(b.position.x-width)
	}

	if b.position.y <= view_radius {
		accel.y = 1 / b.position.y
	} else if b.position.y >= height-view_radius {
		accel.y = -1 / math.Abs(b.position.y-height)
	}
	return accel
}

func (b *Boid) MoveOne() {
	acc := Vector2D{0.0, 0.0}

	average_velocity := Vector2D{0.0, 0.0}
	average_position := Vector2D{0.0, 0.0}
	separation := Vector2D{0.0, 0.0}

	lock.RLock()
	count := 0.0
	for _, boid := range boids {
		if boid.id != b.id {
			dist := boid.position.Dist(b.position)
			if dist < view_radius {
				count++
				average_velocity = average_velocity.Add(boid.velocity)
				average_position = average_position.Add(boid.position)
				separation = separation.Add(b.position.Sub(boid.position).DivV(dist))
			}
		}
	}
	lock.RUnlock()

	lock.Lock()
	if count > 0 {
		average_velocity = average_velocity.DivV(count)
		average_position = average_position.DivV(count)

		alignment_acc := average_velocity.Sub(b.velocity).MulV(acceleration_rate)
		cohesion_acc := average_position.Sub(b.position).MulV(acceleration_rate)
		separation_acc := separation.MulV(acceleration_rate)

		acc = acc.Add(alignment_acc).Add(cohesion_acc).Add(separation_acc)
	}

	avoidance := b.Avoid(screen_width, screen_height)
	// fmt.Println(avoidance)

	b.velocity = b.velocity.Add(acc).Add(avoidance).Clamp(-1, 1)
	b.position = b.position.Add(b.velocity)

	// next := b.position.Add(b.velocity)
	// if next.x >= screen_width || next.x <= 0 {
	// 	b.velocity = Vector2D{-b.velocity.x, b.velocity.y}
	// }

	// if next.y >= screen_height || next.y <= 0 {
	// 	b.velocity = Vector2D{b.velocity.x, -b.velocity.y}
	// }

	// spatialHashGrid.Remove(b.position, b.id)
	// spatialHashGrid.Add(*b)
	lock.Unlock()
}
