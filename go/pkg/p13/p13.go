package p13

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Run() {
	bytes, err := os.ReadFile("../assets/popular-names.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")

	var sb strings.Builder
	sb.Grow(len(bytes))
	r := regexp.MustCompile(`\t`)
	for i, l := range lines {
		if i >= 10 {
			sb.WriteString(strings.Join(lines[i:], "\n"))
			break
		}
		sb.WriteString(string(r.ReplaceAll([]byte(l), []byte(" "))))
	}

	fmt.Println(sb.String())
}
