package main

import "fmt"

func main() {
	v1 := []rune("パトカー")
	v2 := []rune("タクシー")

	var result string
	for i := 0; i < len(v1); i++ {
		result += string(v1[i]) + string(v2[i])
	}
	fmt.Printf("result: %s\n", result)

}
