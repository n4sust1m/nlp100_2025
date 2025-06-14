package main

import (
	"os"

	"github.com/n4sust1m/nlp100_2025/pkg/p00"
	"github.com/n4sust1m/nlp100_2025/pkg/p01"
	"github.com/n4sust1m/nlp100_2025/pkg/p02"
	"github.com/n4sust1m/nlp100_2025/pkg/p03"
	"github.com/n4sust1m/nlp100_2025/pkg/p04"
	"github.com/n4sust1m/nlp100_2025/pkg/p05"
	"github.com/n4sust1m/nlp100_2025/pkg/p06"
	"github.com/n4sust1m/nlp100_2025/pkg/p07"
	"github.com/n4sust1m/nlp100_2025/pkg/p08"
	"github.com/n4sust1m/nlp100_2025/pkg/p09"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("Run with question index\nexample `go run cli/main.go p00`")
	}

	switch args[1] {
	case "p00":
		p00.Run()
	case "p01":
		p01.Run()
	case "p02":
		p02.Run()
	case "p03":
		p03.Run()
	case "p04":
		p04.Run()
	case "p05":
		p05.Run()
	case "p06":
		p06.Run()
	case "p07":
		p07.Run()
	case "p08":
		p08.Run()
	case "p09":
		p09.Run()
	}
}
