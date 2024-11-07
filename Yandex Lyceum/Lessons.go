package main

import (
	"fmt"
	"sort"
)

// Permutations возвращает все перестановки символов в строке input в алфавитном порядке.
func Permutations(input string) []string {
	v, result := []string{}, []string{}
	ger := ""
	liter := 1
	for i := 0; i < len(input); i++ {
		v = append(v, string(input[i]))
	}
	for i := 1; i <= len(v); i++ {
		liter *= i
	}
	for i := 0; i < liter; i++ {
		if len(result) == liter {
			break
		}
		if len(v) == (len(input)*2)-1 {
			v = v[len(input)-1:]
			for u := 0; u < len(v); u++ {
				ger += v[u]
			}
			result = append(result, ger)
			ger = ""
			i = 0
		}
		v = append(v, v[i])
		v[i] = v[i+1]
	}
	sort.Strings(result)
	return result
}
func main() {
	a := "about"
	fmt.Println(Permutations(a))
}
