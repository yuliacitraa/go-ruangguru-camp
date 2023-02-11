package main

import "fmt"

func SchedulableDays(villager [][]int) []int {
	calendar := make(map[int]int)
	for _, avail := range villager {
		for _, day := range avail {
			calendar[day] = calendar[day] + 1
		}
	}
	ppl := len(villager)
	out := make([]int, 0)
	for key, val := range calendar {
		if val == ppl {
			out = append(out, key)
		}
	}
	
	return out // TODO: replace this
}

func main()  {
	data := [][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19},{12, 19}}
	fmt.Println(SchedulableDays(data))
}
