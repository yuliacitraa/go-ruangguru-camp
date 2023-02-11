package main

import (
	"fmt"
)

func PhoneNumberChecker(number string, result *string) {
	var char string

	for i := 0; i < len(number); i++ {
		char = number
	}

	if (char[:2]) == "08" && len(char) >=10 {
		switch char[2:4] {
		case  "11", "12", "13", "14", "15" :
			*result = "Telkomsel"
		case "16", "17", "18", "19" :
			*result = "Indosat"
		case "21", "22", "23" :
			*result = "XL"
		case "27", "28", "29" :
			*result = "Tri"
		case "52", "53" :
			*result = "AS"
		case "81", "82", "83", "84", "85", "86", "87", "88" :
			*result = "Smartfren"
		default :
			*result = "invalid"
		}
	} else if (char[:3]) == "628" && len(char) >= 11 {
		switch char[3:5] {
		case  "11", "12", "13", "14", "15" :
			*result = "Telkomsel"
		case "16", "17", "18", "19" :
			*result = "Indosat"
		case "21", "22", "23" :
			*result = "XL"
		case "27", "28", "29" :
			*result = "Tri"
		case "52", "53" :
			*result = "AS"
		case "81", "82", "83", "84", "85", "86", "87", "88" :
			*result = "Smartfren"
		default :
			*result = "invalid"
		
		}
	} else {
		*result = "invalid"
	}
	// fmt.Println(char[3:5])
}




func main() {
	// bisa digunakan untuk pengujian test case
	var number = "08993456123"
				// "1234567", "invalid",
				// "123", "invalid"
				// "08111111", "invalid"
				// "6281234", "invalid"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)

	// PhoneNumberChecker("6285389898989")

}
