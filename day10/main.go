package main

import (
	"fmt"
	"sort"
	"strconv"

	"aoc2020/shared"
)

func main() {
	part1()

	part2()
}

func part1() {
	numbers, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	sort.Ints(numbers)

	numbers = append([]int{0}, numbers...)

	diffs := map[int]int{}

	for i, n := range numbers {
		if i == len(numbers)-1 {
			break
		}

		d := diff(numbers[i+1], n)
		diffs[d]++
	}

	// for the device
	diffs[3]++

	fmt.Println(diffs[1] * diffs[3])
}

func part2() {
	numbers, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	sort.Ints(numbers)

	acc := map[int]int{0: 1}

	for _, n := range numbers {
		acc[n] = acc[n-1] + acc[n-2] + acc[n-3]
	}

	fmt.Println(acc[numbers[len(numbers)-1]])
}

func diff(h, l int) int {
	return h - l
}

func parse() ([]int, error) {
	lines, err := shared.ReadLines("./input.txt")
	if err != nil {
		return nil, err
	}

	var set []int

	for _, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		set = append(set, n)
	}

	return set, nil
}
