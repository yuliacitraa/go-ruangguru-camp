package main

import "fmt"

func GetTicketPrice (VIP, regular, student, day int) float32 {
	sumTicket := VIP + regular + student
	VIP *= 30
	regular *= 20
	student *= 10
	sumPrice := VIP + regular + student
	var y float32 = float32(sumPrice)
	var x float32

	switch {
		case day % 2 == 1 :
			if sumPrice >= 100 && sumTicket <5 {
				x = 0.15 * y
				
			} 
			if sumPrice >= 100 && sumTicket >=5 {
				x = 0.25 * y
				
			}
		case day % 2 == 0 :
			if sumPrice >= 100 && sumTicket <5 {
				x = 0.1 * y
				
			} 
			if sumPrice >= 100 && sumTicket >=5 {
				x = 0.2 * y
				
			}
	}
	return float32(sumPrice) - x
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
