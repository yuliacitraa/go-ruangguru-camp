package main

import "fmt"

func CountingLetter(text string) int {
	counter:=0
	for _, character := range text {
		letter := string(character)
		if letter == "R" || letter == "r" {
			counter++
		}
		if letter == "S" || letter == "s" {
			counter++
		}
		if letter == "T" || letter == "t" {
			counter++
		}
		if letter == "Z" || letter == "z" {
			counter++
		}

	}
	
	return counter
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Zebra Zig Zag"))
}
