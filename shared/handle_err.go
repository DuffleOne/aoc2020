package shared

import (
	"aoc2020/cher"
	"encoding/json"
	"fmt"
)

func HandleErr(err error) {
	if v, ok := err.(cher.E); ok {
		bytes, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			panic(err)
		}

		fmt.Println("Handled error:")
		fmt.Println(string(bytes))
		panic(err)
	}

	panic(err)
}
