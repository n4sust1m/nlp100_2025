package main

import "fmt"

func main() {
	fmt.Printf("%+v\n", printWithTemplate(12, "気温", 22.4))
}

func printWithTemplate(x int, y string, z float64) string {
	return fmt.Sprintf("%d時の%sは%.1f", x, y, z)
}
