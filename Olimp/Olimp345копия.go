package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, max, min, cent float64
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
	var xx, xn, xt int = int(max), int(min), int(cent)
	if arf && geom {
		fmt.Println("B", xn, xt, xx)
	} else if arf {
		fmt.Println("A", xn, xt, xx)
	} else if geom {
		fmt.Println("G", xn, xt, xx)
	} else {
		fmt.Println("No")
	}
}

func Progressive(a, b, c float64) (bool, bool, []float64) {
	var cent float64
	var max float64 = float64(-100000)
	r := math.Pow(10, 12)
	var min float64 = float64(r)
	arf := false
	geom := false
	nums := []float64{a, b, c}
	for i := 0; i < len(nums); i++ {
		if nums[i] > float64(max) {
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
	if (max-cent == cent-min) && (max-cent >= 0 && cent-min >= 0) {
		arf = true
	}

	if (math.Ceil((max/cent)*100)/100 == cent/min) && math.Ceil((max/cent)*100)/100 >= 1 && cent/min >= 1 {
		geom = true
	}
	masval := []float64{min, cent, max}
	return arf, geom, masval
}
