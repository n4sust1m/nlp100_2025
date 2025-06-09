package main

import "fmt"

func main() {
	v1 := "週刊ascii"
	fmt.Printf("%+v\n", cipher(v1))
	fmt.Printf("%+v\n", cipher(cipher(v1)))
}

func cipher(s string) string {
	result := ""

	for _, c := range s {
		ascii := int(c)

		if int('a') <= ascii && ascii <= int('z') {
			result += string(rune(219 - ascii))
		} else {
			result += string(rune(ascii))
		}
	}

	return result

}
