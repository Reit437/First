package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, max, min, cent int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	arf, geom, masval := Progressive(a, b, c)
	for i := range masval {
		if i == 0 {
			min = masval[i]
		} else if i == 1 {
			cent = masval[i]
		} else {
			max = masval[i]
		}
	}
	if arf && geom {
		fmt.Println("B", min, cent, max)
	} else if arf {
		fmt.Println("A", min, cent, max)
	} else if geom {
		fmt.Println("G", min, cent, max)
	} else {
		fmt.Println("No")
	}
}

func Progressive(a, b, c int) (bool, bool, []int) {
	var cent int
	r := math.Pow(10, 12)
	var min int = int(r)
	var max int = int(r * (-1))
	arf := false
	geom := false
	nums := []int{a, b, c}
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if (nums[i] > min && nums[i] < max) || (min == max) {
			cent = nums[i]
		}
	}
	if (max-cent == cent-min) && (max-cent >= 1 && cent-min >= 1) {
		arf = true
	}
	if (float64(max)/float64(cent) == float64(cent)/float64(min)) && (float64(max)/float64(cent) >= 1 && float64(cent)/float64(min) >= 1) {
		geom = true
	}
	masval := []int{min, cent, max}
	return arf, geom, masval
}
