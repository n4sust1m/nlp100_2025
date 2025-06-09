package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	v1 := "I couldn’t believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

	words := strings.Split(v1, " ")
	result := []string{}

	for _, w := range words {
		if len(w) <= 4 {
			result = append(result, w)
			continue
		}

		// ランダムに並べ替え
		word := []rune(w)
		minIndex := 1
		maxIndex := len(word) - 1
		for i := minIndex; i <= maxIndex; i++ {
			j := minIndex + rand.Intn(maxIndex-minIndex)
			word[i], word[j] = word[j], word[i]
		}
		result = append(result, string(word))
	}

	fmt.Printf("result: %s", strings.Join(result, " "))
}
