package main

import (
	"fmt"
	"sort"
	"strconv"

	"aoc2020/cher"
	"aoc2020/shared"
)

const preamble = 25

func main() {
	num := part1()

	fmt.Println(num)

	part2(num)
}

func part1() int {
	numbers, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	preambleSet := map[int]struct{}{}

	for i := 0; i < preamble; i++ {
		preambleSet[numbers[i]] = struct{}{}
	}

	for i := preamble; i < len(numbers); i++ {
		num := numbers[i]

		found := false
		for n := range preambleSet {
			find := num - n

			if _, ok := preambleSet[find]; ok {
				found = true
				break
			}
		}

		if !found {
			return num
		}

		preambleSet[num] = struct{}{}
		delete(preambleSet, numbers[i-preamble])
	}

	return 0
}

func part2(goal int) {
	numbers, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	first, last := getCombo(numbers, goal)

	if first == 0 || last == 0 {
		shared.HandleErr(cher.New("combo_failed", nil))
	}

	r := numbers[first : last+1]

	sort.Ints(r)

	final := r[0] + r[len(r)-1]

	fmt.Println(final)
}

func getCombo(numbers []int, goal int) (first int, last int) {
ROOT:
	for i1 := range numbers {
		first = i1
		v := goal

		for i2, n2 := range numbers[i1:] {
			last = i1 + i2
			v -= n2

			if v == 0 {
				return first, last
			}

			if v < 0 {
				continue ROOT
			}
		}
	}

	return 0, 0
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
