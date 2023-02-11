package main

import (
	"fmt"
	"strconv"
)

func ReverseData(arr [5]int) [5]int {
	var out [5]int
	var numStr, dummy string

	for i := len(arr) - 1; i >= 0; i-- {
		temp := arr[i]
		// for j := temp - 1; j >= 0; j-- {
		//	fmt.Println(temp)
		//}

		numStr = strconv.Itoa(temp)
		dummy = ""
		for _, char := range numStr {
			dummy = fmt.Sprintf("%c", char) + dummy
		}

		out[4-i], _ = strconv.Atoi(dummy)
	}

	return out // TODO: replace this
}

func main() {
	// var data = [5]int{14,13,12,11,10}
	fmt.Println(ReverseData([5]int{10, 11, 12, 13, 14}))
}
