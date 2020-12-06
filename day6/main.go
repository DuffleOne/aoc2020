package main

import (
	"fmt"
	"strings"

	"aoc2020/shared"
)

func main() {
	part1()
	part2()
}

func part1() {
	lineSets, err := shared.ReadGroups("./input.txt")
	if err != nil {
		panic(err)
	}

	var totalCount int

	for _, ls := range lineSets {
		totalCount += countLineSum(ls)
	}

	fmt.Println(totalCount)
}

func part2() {
	lineSets, err := shared.ReadGroups("./input.txt")
	if err != nil {
		panic(err)
	}

	var totalCount int

	for _, ls := range lineSets {
		totalCount += countLineAll(ls)
	}

	fmt.Println(totalCount)
}

func countLineSum(set []string) int {
	line := strings.Join(set, "")

	m := map[rune]struct{}{}

	for _, l := range line {
		m[l] = struct{}{}
	}

	return len(m)
}

func countLineAll(set []string) int {
	peopleCount := len(set)

	m := map[rune]int{}

	join := strings.Join(set, "")
	for _, l := range join {
		m[l]++
	}

	var t int
	for _, v := range m {
		if v == peopleCount {
			t++
		}
	}

	return t
}
