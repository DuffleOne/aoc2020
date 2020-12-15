package shared

import (
	"strings"
)

func ReadGroups(fileName string) ([][]string, error) {
	var out [][]string

	lines, err := ReadLines(fileName)
	if err != nil {
		return nil, err
	}

	var g []string
	for _, l := range lines {
		if len(l) == 0 {
			out = append(out, g)
			g = []string{}
			continue
		}

		g = append(g, l)
	}

	out = append(out, g)

	return out, nil
}

func ReadGroupsAsLine(fileName string, separator string) ([]string, error) {
	var out []string

	lines, err := ReadLines(fileName)
	if err != nil {
		return nil, err
	}

	var g string
	for _, l := range lines {
		if len(l) == 0 {
			out = append(out, g)
			g = ""
			continue
		}

		g = strings.Join([]string{g, l}, separator)
		g = strings.TrimPrefix(g, separator)
	}

	out = append(out, g)

	return out, nil
}
