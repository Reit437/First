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
	var max float64 = float64(-10000)
	var min float64 = float64(10000)
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
		if (nums[i] > min && nums[i] < max) || min == max {
			cent = nums[i]
		}
	}
	if math.Round((max-cent)*100)/100 == math.Round((cent-min)*100)/100 {
		arf = true
	}
	if max/cent == cent/min {
		geom = true
	}
	masval := []float64{min, cent, max}
	return arf, geom, masval
}
