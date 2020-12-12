package main

import (
	"fmt"
	"regexp"
	"strconv"

	"aoc2020/cher"
	"aoc2020/shared"
)

func main() {
	part1()
	part2()
}

func part1() {
	moves, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	ship := NewShip()

	for _, m := range moves {
		ship.Move(m)
	}

	fmt.Println(ship.ManhattanDistance())
}

func part2() {
	moves, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	ship := NewShip()
	waypoint := NewWaypoint()

	waypoint.S = ship

	for _, m := range moves {
		waypoint.Move(m)
	}

	fmt.Println(waypoint.S.ManhattanDistance())
}

var re = regexp.MustCompile(`^([A-Z])(\d+)$`)

func parse() ([]*Move, error) {
	lines, err := shared.ReadLines("./input.txt")
	if err != nil {
		return nil, err
	}

	set := []*Move{}

	for _, l := range lines {
		if len(l) == 0 {
			break
		}

		match := re.FindStringSubmatch(l)

		if len(match) != 3 {
			return nil, cher.New("invalid_line", cher.M{"line": l})
		}

		n, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, cher.New("invalid_nuber", cher.M{"line": l}, cher.Coerce(err))
		}

		set = append(set, &Move{
			Type:  match[1],
			Value: n,
		})
	}

	return set, nil
}
