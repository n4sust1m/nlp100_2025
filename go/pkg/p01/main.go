package p01

import "fmt"

func Run() {
	v1 := []rune("パタトクカシーー")

	var result string
	for i := 0; i < len(v1); i = i + 2 {
		result += string(v1[i])
	}
	fmt.Printf("result: %s\n", result)
}
