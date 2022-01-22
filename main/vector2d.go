package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

func (v Vector2D) Add(x Vector2D) Vector2D {
	return Vector2D{v.x + x.x, v.y + x.y}
}

func (v Vector2D) Sub(x Vector2D) Vector2D {
	return Vector2D{v.x - x.x, v.y - x.y}
}

func (v Vector2D) Mul(x Vector2D) Vector2D {
	return Vector2D{v.x * x.x, v.y * x.y}
}

func (v Vector2D) Div(x Vector2D) Vector2D {
	if x.x == 0.0 || x.y == 0.0 {
		panic("Division by zero")
	}
	return Vector2D{v.x / x.x, v.y / x.y}
}

func (v Vector2D) AddV(x float64) Vector2D {
	return Vector2D{v.x + x, v.y + x}
}

func (v Vector2D) SubV(x float64) Vector2D {
	return Vector2D{v.x - x, v.y - x}
}

func (v Vector2D) MulV(x float64) Vector2D {
	return Vector2D{v.x * x, v.y * x}
}

func (v Vector2D) DivV(x float64) Vector2D {
	if x == 0.0 {
		panic("Division by zero error")
	}
	return Vector2D{v.x / x, v.y / x}
}

func (v Vector2D) Dist(x Vector2D) float64 {
	return math.Sqrt(math.Pow(v.x-x.x, 2) + math.Pow(v.y-x.y, 2))
}

func (v Vector2D) Clamp(lower, upper float64) Vector2D {
	return Vector2D{
		math.Min(math.Max(v.x, lower), upper),
		math.Min(math.Max(v.y, lower), upper),
	}
}

func (v Vector2D) Normalize() Vector2D {
	length := v.Dist(Vector2D{0, 0})
	return Vector2D{v.x / length, v.y / length}
}
