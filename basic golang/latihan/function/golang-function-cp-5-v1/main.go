package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	min := 9999
    for _, v := range nums {
		if v < min {
			min = v
		}
    }
    return min
}

func FindMax(nums ...int) int {
	max := 0
    for _, v := range nums {
		if v > max {
			max = v
		}
    }
    return max
}

func SumMinMax(nums ...int) int {
	sum := FindMin(nums...) + FindMax(nums...)
	return sum // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(SumMinMax(333, 456, 654, 123, 111, 1000, 1500, 2000, 3000, 1250, 1111))

}
