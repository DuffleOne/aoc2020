package shared

import (
	"bufio"
	"errors"
	"io"
	"os"
)

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	if err != nil {
		return nil, err
	}

	var set []string

	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		if isPrefix {
			panic(errors.New("isPrefix is true"))
		}

		set = append(set, string(line))
	}

	return set, nil
}
