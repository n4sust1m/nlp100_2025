package p05

import (
	"fmt"
	"strings"
)

func Run() {
	v1 := "I am an NLPer"

	fmt.Println("By character")
	fmt.Printf("bi-gram: %+v\n", nGramByChar(2, v1))
	fmt.Printf("tri-gram: %+v\n", nGramByChar(3, v1))

	fmt.Println("By word")
	fmt.Printf("bi-gram: %+v\n", nGramByWord(2, v1))
	fmt.Printf("tri-gram: %+v\n", nGramByWord(3, v1))
}

func nGramByChar(n int, s string) []string {
	result := []string{}

	for i := 0; i < len(s); i++ {
		end := i + n
		if end > len(s) {
			break
		}
		result = append(result, s[i:end])
	}

	return result
}

func nGramByWord(n int, s string) [][]string {
	words := strings.Split(s, " ")

	result := [][]string{}
	for i := 0; i < len(words); i++ {
		end := i + n
		if end > len(words) {
			break
		}
		result = append(result, words[i:end])
	}

	return result
}
