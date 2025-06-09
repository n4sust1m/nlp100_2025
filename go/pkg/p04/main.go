package main

import (
	"fmt"
	"regexp"
	"slices"
)

func main() {
	v1 := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	specifiedIndex := []int{
		1, 5, 6, 7, 8, 9, 15, 16, 19,
	}

	splitted := regexp.MustCompile(`[\s\,\.]+`).Split(v1, -1)

	result := map[string]int{}
	for i, v := range splitted {
		if len(v) == 0 {
			continue
		}

		if slices.Index(specifiedIndex, i) != -1 {
			result[string(v[0])] = i
		} else {
			symbol := v[:2]
			result[symbol] = i
		}
	}

	fmt.Printf("result: %+v\n", result)
}
