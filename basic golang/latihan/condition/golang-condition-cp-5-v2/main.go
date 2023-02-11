package main

import "fmt"

func TicketPlayground(height, age int) int {
	if age > 12 {
		return 100000
	} else if height >= 160 || age == 12 {
		return 60000
	} else if  height >= 150 || age >=10 && age <=11 {
		return 40000
	} else if  height >= 135 || age >=8 && age <=9 {
		return 25000
	} else if height >= 120 || age >= 5 && age <=7 {
		return 15000
	} else {
		return -1
	}			
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(165, 10))
}
