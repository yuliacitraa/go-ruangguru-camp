package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Reverse(str string) string {
	var out string
	for i := len(str) - 1; i >= 0; i-- {
		pass := string(str[i])
		// if pass == " " {
		// 	continue
		// }
		out += pass
	}
	return out
}

func Generate(str string) string {
	str = Reverse(str)
	var out string
	for _, v := range str {
		var char string
		char = fmt.Sprintf("%c", v)

		switch char {
		case "A":
			char = "E"
		case "a":
			char = "e"
		case "I":
			char = "O"
		case "i":
			char = "o"
		case "U":
			char = "A"
		case "u":
			char = "a"
		case "E":
			char = "I"
		case "e":
			char = "i"
		case "O":
			char = "U"
		case "o":
			char = "u"
		case " " :
			continue
		}

		lower := false
		if unicode.IsUpper(v) {
			lower = true
		}

		if !lower {
			char = strings.ToUpper(char)
		}else {
			char = strings.ToLower(char)
		}
		out = out + char
	}

	// var final string

	return out
}

func CheckPassword(str string) string {
	length := len(str)

	symb := 0
	for _, v := range str {
		switch  {
		case unicode.IsLetter(v):
			continue
		case unicode.IsNumber(v):
			continue
		default :
			symb++			
		}
	}

	switch {
		case length < 7 :
			return "sangat lemah"
		case length >= 7 && symb == 0 :
			return "lemah"
		case length < 14 && symb > 0 :
			return "sedang"
		case length >= 14 && symb > 0 :
			return "kuat"
	}
	
	return "" // TODO: replace this
}

func PasswordGenerator(base string) (string, string) {
	// reverse := Reverse(base)
	generate := Generate(base)
	check := CheckPassword(generate)
	return generate, check // TODO: replace this
}

func main() {
	data := "SEMANGAT PAGI" // bisa digunakan untuk melakukan debug

	// res, check := PasswordGenerator(data)

	// fmt.Println(res)
	// fmt.Println(check)

	fmt.Println(Reverse(data))
}
