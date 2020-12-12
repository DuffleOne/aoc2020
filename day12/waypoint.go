package main

import (
	"image"
	"math"
)

type Waypoint struct {
	P image.Point
	S *Ship
}

func NewWaypoint() *Waypoint {
	return &Waypoint{
		P: image.Point{10, 1},
		S: nil,
	}
}

func (w *Waypoint) Move(m *Move) {
	switch m.Type {
	case "F":
		for i := 0; i < m.Value; i++ {
			w.S.P = w.S.P.Add(w.P)
		}
	case "R":
		r := m.Value / 90
		for i := 0; i < r; i++ {
			w.P = image.Point{w.P.Y, -w.P.X}
		}
	case "L":
		r := m.Value / 90
		for i := 0; i < r; i++ {
			w.P = image.Point{-w.P.Y, w.P.X}
		}
	case "N":
		w.P = w.P.Add(image.Point{0, m.Value})
	case "S":
		w.P = w.P.Add(image.Point{0, -m.Value})
	case "E":
		w.P = w.P.Add(image.Point{m.Value, 0})
	case "W":
		w.P = w.P.Add(image.Point{-m.Value, 0})
	}
}

func (w Waypoint) ManhattanDistance() int {
	x := float64(w.P.X)
	y := float64(w.P.Y)

	return int(math.Abs(x) + math.Abs(y))
}
