package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	a := strings.Split(str, " ")
	var out string
	for _, v := range a {
		var temp string
		var isCapital bool
		for i := len(v) - 1; i >= 0; i-- {
			temp += strings.ToLower(string(v[i]))
			if i == 0 && unicode.IsUpper(rune(v[i])) {
				isCapital = true
			}
		}

		//cek apabila huruf pertama dalam sebuah kata adalah capital
		if isCapital {
			temp = strings.Title(temp)
		}
		out += temp
		// // untuk cek apakah perlu ditambahkan spasi di akhir kata
		// if idx != len(v)-1 || len(v) == 1 {
		// 	out += " "
		// }
		out += " "
	}
	return out[:len(out)-1]

	// var out string

	// for _, v := range strings.Split(str, " ") {
	// 	upper := false
	// 	if unicode.IsUpper(rune(v[0])) {
	// 		upper = true
	// 	} 

	// 	for i := len(v)-1; i >= 0; i-- {
	// 		out+=string(v[i])
	// 	}
	// 	out += " "

	// 	out = strings.ToLower(out)

	// 	if upper == true {
	// 		out = strings.Title(out)
	// 	} else {
	// 		out = strings.ToLower(out)
	// 	}
	// }
	// return out

	// a := strings.Split(str, " ")
	// var out string
	// for idx, v := range a {
	// 	var temp string
	// 	for i := len(v) - 1; i >= 0; i-- {
	// 		//upper case huruf pertama dlm sebuah kata sisanya lower case
	// 		if i == len(v)-1 {
	// 			temp += strings.ToUpper(string(v[i]))
	// 		} else {
	// 			temp += strings.ToLower(string(v[i]))
	// 		}
	// 	}

	// 	out += temp
	// 	// untuk cek apakah perlu ditambahkan spasi di akhir kata
	// 	if idx != len(v)-1 {
	// 		out += " "
	// 	}
	// }
	// return out

	// sliceStr := strings.Split(str, " ")
	// var temp []string 
	// var out string 

	// for _, v := range sliceStr {
	// 	for i := len(v)-1 ; i >= 0 ; i-- {
	// 		// switch {
	// 		// case i == 0 && len(v) > 1 :
	// 		// 	out += strings.ToLower(string(v[i]))
	// 		// // case i == len(v)-1 || len(v) ==1:
	// 		// // 	out += strings.ToUpper(string(v[i]))
	// 		// default :
	// 		// 	out += string(v[i])
	// 		// }
	// 		 out += string(v[i])
	// 	}
	// 	temp = append(temp, out)
	// 	out=""
	// 	// out += " "
	// }

	// var res string
	// for _, v := range temp {
	// 	for i := 0; i < len(v); i++ {
			
	// 	}
	// 	if unicode.IsUpper(rune(v[0])) {
	// 		res += string(v[0])
	// 	}
	// 	if unicode.IsUpper(rune(v[len(v)-1])) {
	// 		res += strings.ToUpper(string(v[0]))
	// 	}
		
	// }



	// return res

	// slipt str
	// reverse semua str (reverse loop)
	// cek upper dan lower char
	// into rules
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu")) //Uka Gnayas Ubi
	fmt.Println(ReverseWord("A bird fly to Germany and got a worm")) //A drib ylf ot Ynamreg dna tog a mrow
	fmt.Println(ReverseWord("ini terlalu mudah")) //ini ulalret hamud
	fmt.Println(ReverseWord("KITA SELALU BERSAMA")) //Atik Ulales Amasreb
}
