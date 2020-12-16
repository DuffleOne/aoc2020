package main

import (
	"fmt"

	"aoc2020/shared"
)

func main() {
	numbers, err := shared.ReadNumbers("input.txt", shared.Comma)
	if err != nil {
		shared.HandleErr(err)
	}

	fmt.Println(RunGame(numbers, 2020))

	fmt.Println(RunGame(numbers, 30000000))
}

func RunGame(numbers []int, iterations int) int {
	spoken, turn, last := map[int]int{}, 0, 0

	for _, n := range numbers {
		if turn > 0 {
			spoken[last] = turn
		}

		last = n
		turn++
	}

	for turn < iterations {
		if t, ok := spoken[last]; ok {
			spoken[last] = turn
			last = turn - t
		} else {
			spoken[last] = turn
			last = 0
		}

		turn++
	}

	return last
}
