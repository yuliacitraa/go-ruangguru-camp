package main

import (
	"fmt"
	"strconv"
)

func DateFormat(day, month, year int) string {
	bulan := strconv.Itoa(month)
	var out string

	switch bulan {
	case "1" :
		bulan = "January" 
	case "2" :
		bulan = "February" 
	case "3" :
		bulan = "March"
	case "4" :
		bulan = "April"
	case "5" :
		bulan = "May"
	case "6" :
		bulan = "June"
	case "7" :
		bulan = "July"
	case "8" :
		bulan = "August"
	case "9" :
		bulan = "September"
	case "10" :
		bulan = "October"
	case "11" :
		bulan = "November"
	case "12" :
		bulan = "December"
	}
	
	out = fmt.Sprintf("%02d-%s-%d", day, bulan, year)

	return out
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 2, 2012))
}
