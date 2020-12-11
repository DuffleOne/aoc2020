package main

import (
	"fmt"
	"image"

	"aoc2020/shared"
)

const (
	floor     = '.'
	empty     = 'L'
	occcupied = '#'
)

var pointMap = []image.Point{
	{-1, -1}, // top left
	{-1, 0},  // left
	{-1, 1},  // bottom left
	{0, -1},  // top
	{0, 1},   // bottom
	{1, -1},  // top right
	{1, 0},   // right
	{1, 1},   // bottom right
}

func main() {
	seats, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	fmt.Println(runRules(seats, 4))
}

func parse() (map[image.Point]rune, error) {
	lines, err := shared.ReadLines("./input.txt")
	if err != nil {
		return nil, err
	}

	seats := map[image.Point]rune{}

	for rowIndex, l := range lines {
		for colIndex, m := range l {
			seats[image.Point{rowIndex, colIndex}] = m
		}
	}

	return seats, nil
}

func runRules(seats map[image.Point]rune, radius int) (out int) {
	for {
		totalOccupied := 0
		// you shouldn't change the values as you iterate over
		// so we'll make a new set and replicate seats each iteration
		iSeats := map[image.Point]rune{}

		for pos, seatType := range seats {
			var seatCount int

			// get count of occupied adjacent seats
			for _, coord := range pointMap {
				if seats[getAdjacent(pos, coord)] == occcupied {
					seatCount++
				}
			}

			// apply rules
			switch {
			case seatType == empty && seatCount == 0:
				seatType = occcupied
			case seatType == occcupied && seatCount >= radius:
				seatType = empty
			}

			if seatType == occcupied {
				totalOccupied++
			}

			iSeats[pos] = seatType
		}

		// if none have changed since the last iteration we're done
		if totalOccupied == out {
			return out
		}

		// setup the next round
		out = totalOccupied
		seats = iSeats
	}
}

func getAdjacent(p1, p2 image.Point) image.Point {
	return p1.Add(p2)
}
