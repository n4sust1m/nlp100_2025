package main

import (
	"fmt"
	"regexp"
)

func main() {
	v1 := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."

	splitted := regexp.MustCompile(`[\s\,\.]+`).Split(v1, -1)
	var result []int
	for _, word := range splitted {
		if len(word) == 0 {
			continue
		}
		result = append(result, len(word))
	}
	fmt.Printf("result: %+v\n", result)
}
