package main

import "fmt"

func CountingNumber(n int) float64 {
	var i, total float64
	j := float64(n)

    for i=1;i<=j;i+=0.5 {
		total += i
	}

	return total
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
