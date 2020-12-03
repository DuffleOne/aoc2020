package main

import (
	"fmt"
	"strconv"

	"aoc2020/shared"
)

const goal = 2020

func main() {
	numbers, err := readFile("./input.txt")
	if err != nil {
		shared.HandleErr(err)
	}

	for _, n := range numbers {
		for _, m := range numbers {
			for _, o := range numbers {
				if numbersAreSame(m, n, o) {
					continue
				}

				if n+m+o == goal {
					fmt.Printf("Found %d + %d + %d = %d. Sum = %d\n", n, m, o, goal, n*m*o)
					return
				}
			}
		}
	}
}

func numbersAreSame(set ...int) bool {
	t := map[int]struct{}{}

	for _, n := range set {
		if _, ok := t[n]; ok {
			return true
		}

		t[n] = struct{}{}
	}

	return false
}

func readFile(fileName string) ([]int, error) {
	lines, err := shared.ReadLines(fileName)
	if err != nil {
		return nil, err
	}

	var out []int

	for _, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		out = append(out, n)
	}

	return out, nil
}
