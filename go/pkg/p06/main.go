package p06

import "fmt"

func Run() {
	v1 := "paraparaparadise"
	v2 := "paragraph"

	x := biGramByChar(v1)
	y := biGramByChar(v2)

	fmt.Printf("X: %+v\n", *x)
	fmt.Printf("Y: %+v\n", *y)
	fmt.Printf("sum: %+v\n", *sum(x, y))
	fmt.Printf("prod: %+v\n", *prod(x, y))
	fmt.Printf("diff: %+v\n", *diff(x, y))

}

type Set = map[string]any // set の代替データ構造

func biGramByChar(s string) *[]string {
	result := []string{}

	for i := 0; i < len(s); i++ {
		end := i + 2
		if end > len(s) {
			break
		}
		result = append(result, s[i:end])
	}

	return &result
}

func sum(x *[]string, y *[]string) *[]string {
	set := Set{}

	for _, v := range [][]string{
		*x,
		*y,
	} {
		for _, k := range v {
			set[k] = 1
		}
	}

	result := []string{}
	for k, _ := range set {
		result = append(result, k)
	}

	return &result
}

func prod(x *[]string, y *[]string) *[]string {
	set := Set{}

	for _, v := range *x {
		set[v] = 1
	}
	for _, v := range *y {
		_, ok := set[v]
		if !ok {
			delete(set, v)
		}
	}

	result := []string{}
	for k, _ := range set {
		result = append(result, k)
	}

	return &result
}

func diff(x *[]string, y *[]string) *[]string {
	set := Set{}

	for _, v := range *x {
		set[v] = 1
	}
	for _, v := range *y {
		_, ok := set[v]
		if ok {
			delete(set, v)
		}
	}

	result := []string{}
	for k, _ := range set {
		result = append(result, k)
	}

	return &result
}
