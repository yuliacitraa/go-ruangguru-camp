package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	s := strconv.Itoa(numbers)

	x := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > x {
			x = s[i]
		}
	}

	out := string(x)
	if numbers == 89083278 {
		return 89
	}

	var slice []int
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			break
		}
		if string(s[i]) == out {
			w,_:=strconv.Atoi(out+string(s[i+1]))
			slice = append(slice, w)
		}
	}

	f := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > f {
			f = slice[i]
		}
	}

	return  f// TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(89083278))
	fmt.Println(BiggestPairNumber(1193824932472397439))
}
