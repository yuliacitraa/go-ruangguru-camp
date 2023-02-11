package main

import "fmt"

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var out string
	for i := len(str)-1; i >= 0; i-- {
		switch {
			case i==len(str)-1 :
				out += string(str[i])
			case string(str[i]) == " ":
				out += string(str[i])
			case string(str[i+1]) == " ":
				out += string(str[i])
			default:
				out += "_" + string(str[i])
		}
	}
	return out
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
