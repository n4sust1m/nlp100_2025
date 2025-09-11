package p11

import (
	"fmt"
	"os"
	"strings"
)

const N int = 10

func Run() {
	r, err := getInitialCharacters("../assets/popular-names.txt")
	if err != nil {
		panic(fmt.Sprintf("err: %+v", err))
	}

	fmt.Println(*r)
}

func getInitialCharacters(fileName string) (*string, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")

	var sb strings.Builder
	for _, l := range lines {
		limit := min(len(l), N)
		sb.WriteString(l[:limit] + "\n")
	}
	result := sb.String()
	return &result, nil
}
