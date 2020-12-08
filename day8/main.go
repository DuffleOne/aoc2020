package main

import (
	"fmt"
	"regexp"
	"strconv"

	"aoc2020/cher"
	"aoc2020/shared"
)

var instructionRegex = regexp.MustCompile(`^([a-z]{3})\s([+|-]\d+)$`)

func main() {
	err := part1()
	shared.HandleErr(err)

	err = part2()
	shared.HandleErr(err)
}

func part1() error {
	comp, err := parse(false)
	if err != nil {
		return err
	}

	err = comp.Run()
	_, ok := err.(cher.E)
	if !ok {
		return err
	}

	fmt.Println(comp.Accumulator)

	return nil
}

func part2() error {
	comp, err := parse(false)
	if err != nil {
		return err
	}

	for _, is := range comp.Instructions {
		if is.Operation == "nop" {
			continue
		}

		is.Swap()
		err = comp.Run()
		if err == nil {
			fmt.Println(comp.Accumulator)
			break
		}
		is.Swap()
	}

	return nil
}

func parse(debug bool) (*InstructionSet, error) {
	lines, err := shared.ReadLines("./input.txt")
	if err != nil {
		return nil, err
	}

	comp := InstructionSet{
		Debug: debug,
	}

	is := []*Instruction{}

	for _, line := range lines {
		if !instructionRegex.MatchString(line) {
			return nil, cher.New("invalid_line", cher.M{"line": line})
		}

		matches := instructionRegex.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			n, err := strconv.Atoi(m[2])
			if err != nil {
				return nil, err
			}

			is = append(is, &Instruction{
				Operation: m[1],
				Value:     n,
			})

		}
	}

	comp.Instructions = is

	return &comp, nil
}

type InstructionSet struct {
	Accumulator  int
	Debug        bool
	Instructions []*Instruction
}

type Instruction struct {
	Operation string
	Swapped   bool
	Value     int
}

func (i *Instruction) Swap() {
	i.Swapped = !i.Swapped
	switch i.Operation {
	case "nop":
		i.Operation = "jmp"
	case "jmp":
		i.Operation = "nop"
	}
}

func (c *InstructionSet) Run() error {
	c.Accumulator = 0
	hasRan := map[int]struct{}{}

	i := 0
	for {
		o := i

		ins := c.Instructions[i]

		if _, ok := hasRan[i]; ok {
			return cher.New("infinite_loop", nil)
		}

		switch ins.Operation {
		case "nop":
			i++
			// do nothing
		case "acc":
			c.Accumulator += ins.Value
			i++
		case "jmp":
			i += ins.Value
		}

		if c.Debug {
			fmt.Printf("%d: %s %d | %d\n", o, ins.Operation, ins.Value, c.Accumulator)
		}

		hasRan[o] = struct{}{}

		if i >= len(c.Instructions) {
			break
		}
	}

	return nil
}
