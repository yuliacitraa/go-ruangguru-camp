package main

import "fmt"

func SlurredTalk(words *string) {
	var out string

	for i := 0; i < len(*words); i++ {
		switch (*words)[i] {
			case 'S', 'R', 'Z' :
				out += "L"
				continue
			case 's', 'r', 'z' :
				out += "l"
				continue	
		}
		out += string((*words)[i])
	}
	
	*words = out
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "saya Steven"
	SlurredTalk(&words)
	fmt.Println(words)
}
