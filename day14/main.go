package main

import (
	"fmt"
	"regexp"
	"strconv"

	"aoc2020/cher"
	"aoc2020/shared"
)

func main() {
	instructions, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	memory := map[int]uint64{}

	for _, i := range instructions {
		for _, o := range i.Operations {
			memory[o.Address] = apply(i.Mask, o.Value)
		}
	}

	var total uint64 = 0
	for _, v := range memory {
		total += v
	}

	fmt.Println(total)
}

func apply(mask string, value uint64) uint64 {
	for i, s := range mask {
		switch s {
		case 'X':
		case '0':
			value &= ^(1 << (35 - i))
		case '1':
			value |= (1 << (35 - i))
		}
	}

	return value
}

var maskRegex = regexp.MustCompile(`^mask = ([X01]+)$`)
var operationRegex = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

func parse() ([]*Instruction, error) {
	lines, err := shared.ReadLines("./input.txt")
	if err != nil {
		return nil, err
	}

	set := []*Instruction{}

	var ins *Instruction

	for len(lines) > 0 {
		currLine := lines[0]

		maskMatch := maskRegex.FindStringSubmatch(currLine)
		operationMatch := operationRegex.FindStringSubmatch(currLine)

		switch {
		case len(maskMatch) == 2:
			if ins != nil {
				set = append(set, ins)
			}
			ins = &Instruction{
				Mask: maskMatch[1],
			}
		case len(operationMatch) == 3:
			addrNum, err := strconv.Atoi(operationMatch[1])
			if err != nil {
				return nil, err
			}
			valueNum, err := strconv.Atoi(operationMatch[2])
			if err != nil {
				return nil, err
			}

			ins.Operations = append(ins.Operations, Operation{
				Address: addrNum,
				Value:   uint64(valueNum),
			})
		default:
			return nil, cher.New("invalid_line", cher.M{"line": currLine})
		}

		lines = lines[1:]
	}

	set = append(set, ins)

	return set, nil
}

type Instruction struct {
	Mask       string
	Operations []Operation
}

type Operation struct {
	Address int
	Value   uint64
}
