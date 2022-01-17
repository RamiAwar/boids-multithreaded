package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	a := Vector2D{x: 1.0, y: 1.0}
	b := Vector2D{x: -1.0, y: 4.0}

	a.Add(b)

	got := a.x
	want := 0.0
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}

	got = a.y
	want = 5.0
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}
}

func TestAddChained(t *testing.T) {
	a := Vector2D{x: 1.0, y: 1.0}
	b := Vector2D{x: -1.0, y: 4.0}

	a.Add(b).Add(b)

	got := a.x
	want := -1.0
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}

	got = a.y
	want = 9.0
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}
}

func TestSub(t *testing.T) {
	a := Vector2D{x: 1.0, y: 1.0}
	b := Vector2D{x: -1.0, y: 4.0}

	a.Sub(b)

	got := a.x
	want := 2.0
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}

	got = a.y
	want = -3.0
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}
}

func TestMul(t *testing.T) {
	a := Vector2D{x: 1.0, y: 1.0}
	b := Vector2D{x: -1.0, y: 4.0}

	a.Mul(b)

	got := a.x
	want := -1.0
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}

	got = a.y
	want = 4.0
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}
}

func TestDiv(t *testing.T) {
	a := Vector2D{x: 1.0, y: 1.0}
	b := Vector2D{x: -1.0, y: 4.0}

	a.Div(b)

	got := a.x
	want := -1.0
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}

	got = a.y
	want = 0.25
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}
}

func TestDivByZero(t *testing.T) {
	a := Vector2D{x: 1.0, y: 1.0}
	b := Vector2D{x: 0.0, y: 0.0}

	// Expect a panic
	defer func() { recover() }()
	a.Div(b)

	t.Errorf("Division by zero error did not cause a panic")
}

func TestDivVByZero(t *testing.T) {
	a := Vector2D{x: 1.0, y: 1.0}
	c := 0.0

	// Expect a panic
	defer func() { recover() }()
	a.DivV(c)

	t.Errorf("Division by zero error did not cause a panic")
}

func TestDist(t *testing.T) {
	a := Vector2D{x: 1.0, y: 1.0}
	origin := Vector2D{x: 0.0, y: 0.0}

	want := math.Sqrt(2)
	got := a.Dist(origin)
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}
}

func TestClamp(t *testing.T) {
	a := Vector2D{-1.2, 1.2}
	a.Clamp(0.0, 1.0)

	want := 0.0
	got := a.x
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}

	want = 1.0
	got = a.y
	if got != want {
		t.Errorf("Expected %f, got %f", want, got)
	}

}
