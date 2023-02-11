package main

import (
	"fmt"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowel, consonant := 0, 0
	var isTrue bool
	
	for _, v := range str {
		if 65 <= v && v <= 90 || 97 <= v && v <= 122 {		
			switch string(v) {
				case "a","i","u","e","o" :
					vowel++
				case "A","I","U","E","O" :
					vowel++
				default :
					consonant++
			}
			
			if vowel>0 && consonant>0 {
				isTrue = false
			} else {
				isTrue = true
			}
		}
	}
	return vowel, consonant, isTrue // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("SEMANGAT PAGI, itu kata orang yang baru datang ke rumahku"))
	fmt.Println(CountVowelConsonant("kopi"))
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))
	fmt.Println(CountVowelConsonant("HALO DUNIA INDAH YANG BAIK"))

}
