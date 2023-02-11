package main

import "fmt"

func CountProfit(data [][][2]int) []int { 
	res := make(map[int]int)

	for _, v := range data {
		for i, j := range v {
			profit := j[0] - j[1]
			res[i] = res[i] + profit

		}
	}
	out := make([]int, len(res))
	for k, v := range res {
		out[k] = v	
	}
	return out // TODO: replace this
}

func main()  {
	fmt.Println(CountProfit([][][2]int{{ {1000, 500}, {500, 200}},{{ 1200, 200},{1000, 800}},{{ 500, 100},{700, 100}   }}))
}


// func main()  {
// 	var numbers = [][][2]int{{{3, 2}, {3, 4}, {3, 5}, {3, 6}}}
// 	fmt.Println(numbers)
// }