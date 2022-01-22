package main

import (
	"fmt"
	"math"
)

type SpatialHash struct {
	CellSize float64
	Cells    map[SpatialHashKey][]Boid
}

type SpatialHashKey struct {
	x, y int
}

func NewSpatialHash(cellSize float64) *SpatialHash {
	cells := make(map[SpatialHashKey][]Boid)
	return &SpatialHash{cellSize, cells}
}

func (h *SpatialHash) KeyForPoint(point Vector2D) SpatialHashKey {
	x := int(math.Floor(point.x / h.CellSize))
	y := int(math.Floor(point.y / h.CellSize))
	return SpatialHashKey{x, y}
}

func (h *SpatialHash) Add(boid Boid) {
	key := h.KeyForPoint(boid.position)
	h.Cells[key] = append(h.Cells[key], boid)
}

func (h *SpatialHash) Remove(point Vector2D, id int) {
	key := h.KeyForPoint(point)
	for i, cell := range h.Cells[key] {
		if cell.id == id {
			// Remove element from slice
			n := len(h.Cells[key])
			h.Cells[key][i] = h.Cells[key][n-1]
			h.Cells[key] = h.Cells[key][:n-1]
			return
		}
	}
	fmt.Println("Failed to remove")
}

func (h *SpatialHash) Nearby(point Vector2D, radius int) []Boid {
	var result []Boid
	key := h.KeyForPoint(point)
	for dx := -radius; dx <= radius; dx++ {
		for dy := -radius; dy <= radius; dy++ {
			k := SpatialHashKey{key.x + dx, key.y + dy}
			result = append(result, h.Cells[k]...)
		}
	}
	return result
}
