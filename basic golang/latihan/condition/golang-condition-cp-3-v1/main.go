package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	average := (math + science + english +indonesia) / 4
	// switch average {
	// 	case average 100 :
	// 		return "sempurna"
	// }
	if average == 100 {
		return "Sempurna"
	} else if average >= 90 {
		return "Sangat Baik"
	} else if average >= 80 {
		return "Baik"
	} else if average >= 70 {
		return "Cukup"
	} else if average >= 60 {
		return "Kurang"
	} else if average < 60 {
		return "Sangat kurang"
	} else {
		return ""
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
