package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"aoc2020/cher"
	"aoc2020/shared"
)

func main() {
	part1()
	part2()
}

func part1() {
	time, lines, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	ans := math.MaxInt64

	for _, line := range lines {
		if line == -1 {
			continue
		}

		if line-time%line < ans-time%ans {
			ans = line
		}
	}

	fmt.Println(ans * (ans - time%ans))
}

func part2() {
	_, lines, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	ans := 0
	step := 1

	for i, line := range lines {
		if line == -1 {
			continue
		}

		for (ans+i)%line != 0 {
			ans += step
		}

		step *= line
	}

	fmt.Println(ans)
}

func parse() (int, []int, error) {
	lines, err := shared.ReadLines("./input.txt")
	if err != nil {
		return 0, nil, err
	}

	if len(lines) != 2 {
		return 0, nil, cher.New("invalid_input", nil)
	}

	number, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, nil, err
	}

	set := []int{}

	for _, s := range strings.Split(lines[1], ",") {
		if s == "x" {
			set = append(set, -1)
			continue
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, nil, err
		}

		set = append(set, n)
	}

	return number, set, nil
}
