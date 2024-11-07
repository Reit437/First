package main

import (
	"fmt"
	"sort"
)

func main() {
	stringSlice := []string{"cab", "abc", "acb", "YaLyceum", "Alpha", "Grin", "Delta"}
	sort.Strings(stringSlice) // Сортировка уже для строк
	fmt.Println(stringSlice)  // [Alpha Bravo Delta Go Gopher Grin YaLyceum]
}
