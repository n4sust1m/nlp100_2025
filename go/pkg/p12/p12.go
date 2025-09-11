package p12

import (
	"fmt"
	"os"
	"strings"
)

const N int = 10

func Run() {
	r, err := getFinalCharacters("../assets/popular-names.txt")
	if err != nil {
		panic(fmt.Sprintf("err: %+v", err))
	}

	fmt.Println(*r)
}

func getFinalCharacters(fileName string) (*string, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")

	var sb strings.Builder
	for _, l := range lines {
		limit := max(len(l)-N, 0)
		sb.WriteString(l[limit:] + "\n")
	}
	result := sb.String()
	return &result, nil
}
