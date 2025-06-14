package p10

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	count, err := countLinesInFile("../assets/popular-names.txt")
	if err != nil {
		panic(fmt.Sprintf("err: %+v", err))
	}

	fmt.Printf("count: %d\n", count)
}

func countLinesInFile(fileName string) (int, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	lines := string(bytes)
	line := strings.Split(lines, "\n")
	return len(line), nil
}
