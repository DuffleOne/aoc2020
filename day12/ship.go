package main

import (
	"image"
	"math"
)

type Direction int

const (
	East Direction = iota
	South
	West
	North
)

var directions = map[Direction]string{
	0: "E",
	1: "S",
	2: "W",
	3: "N",
}

type Ship struct {
	P   image.Point
	Dir Direction
}

func NewShip() *Ship {
	return &Ship{
		P:   image.Point{0, 0},
		Dir: 0,
	}
}

func (s *Ship) Move(m *Move) {
AGAIN:
	switch m.Type {
	case "F":
		m.Type = directions[s.Dir]
		goto AGAIN
	case "R":
		s.Dir += Direction(m.Value / 90)
		if s.Dir > 3 {
			s.Dir -= 4
		}
	case "L":
		s.Dir -= Direction(m.Value / 90)
		if s.Dir < 0 {
			s.Dir += 4
		}
	case "N":
		s.P = s.P.Add(image.Point{0, m.Value})
	case "S":
		s.P = s.P.Add(image.Point{0, -m.Value})
	case "E":
		s.P = s.P.Add(image.Point{m.Value, 0})
	case "W":
		s.P = s.P.Add(image.Point{-m.Value, 0})
	}
}

func (s Ship) ManhattanDistance() int {
	x := float64(s.P.X)
	y := float64(s.P.Y)

	return int(math.Abs(x) + math.Abs(y))
}

type Move struct {
	Type  string
	Value int
}
