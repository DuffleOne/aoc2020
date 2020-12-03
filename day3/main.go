package main

import (
	"strings"

	"aoc2020/cher"
	"aoc2020/shared"

	"github.com/kr/pretty"
)

func main() {
	terrain, err := parseInput()
	if err != nil {
		panic(err)
	}

	var patterns = [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	var treeCounts []int

	for _, pattern := range patterns {
		var treeCount int
		for i := 0; i < len(terrain.Map)-1; i++ {
			terrain.Navigate(pattern[0], pattern[1])
			_, _, tree := terrain.GetPosition()

			if tree == -1 {
				break
			}

			treeCount += tree
		}

		terrain.ResetPosition()
		treeCounts = append(treeCounts, treeCount)
	}

	var acc int = 1

	for _, n := range treeCounts {
		acc *= n
	}

	pretty.Println(acc)
}

func parseInput() (*Terrain, error) {
	lines, err := shared.ReadLines("./input.txt")
	if err != nil {
		return nil, err
	}

	var rows [][]int

	for _, l := range lines {
		set := strings.Split(l, "")
		var row []int

		for _, char := range set {
			switch char {
			case ".":
				row = append(row, 0)
			case "#":
				row = append(row, 1)
			default:
				return nil, cher.New("unrecognisd_character", cher.M{"char": char})
			}
		}

		rows = append(rows, row)
	}

	return &Terrain{
		X:   0,
		Y:   0,
		Map: rows,
	}, nil
}
