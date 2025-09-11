package p02

import "fmt"

func Run() {
	v1 := "stressed"

	var result string
	for i := len(v1) - 1; i >= 0; i-- {
		result += string(v1[i])
	}
	fmt.Printf("result: %s\n", result)
}
