package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	var hasil string

	for _, val := range data {
		if strings.Contains(val, input) {
			hasil = hasil + val + ","
			// if idx == len(data) - 1 {
			// 	continue
			// }
			
		}
	}
	return hasil[:len(hasil)-1] // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
