package p14

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
	r := regexp.MustCompile(`\s`)
	for i, l := range lines {
		if i >= 10 {
			sb.WriteString(strings.Join(lines[i:], "\n"))
			break
		}

		c := r.Split(l, -1)[0]
		sb.WriteString(c + "\n")
	}

	fmt.Println(sb.String())
}
