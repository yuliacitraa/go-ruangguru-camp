package main

import (
	"fmt"
)

func SchedulableDays(date1 []int, date2 []int) []int {	
	free := []int{}

	// for _, v := range date2 {
	// 	for _, v2 := range date1 {
	// 		if v == v2 {
	// 			free = append(free, date2[v])
	// 		}
	// 	}		
	// }

	for i := 0; i < len(date1); i++ {
		for j := 0; j < len(date2); j++ {
			if date1[i] == date2[j] {
	        	free = append(free, date1[i])
	    	}
		}
	}
	return free
}

func main()  {
	imam := []int{11, 12, 13, 14, 15}
	permana := []int{5, 10, 12, 13, 20, 21}

	fmt.Println(SchedulableDays(imam, permana))

}
