package main

import (
	"fmt"
)

func Add(a, b int) int {
	return 0
}

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func findAvg(grades ...int) float64 {
	// rumus : total jumlah nilai / banyak data
	if len(grades) == 0 {
		return 0
	}

	var temp int
	for _, v := range grades {
		temp += v
	}

	return float64(temp) / float64(len(grades))
}

func findMin(grades ...int) int {
	if len(grades) == 0 {
		return 0
	}

	temp := grades[0]

	for i := 1; i < len(grades); i++ {
		if grades[i] <= temp {
			temp = grades[i]
		}
	}

	return temp
}

func findMax(grades ...int) int {

	temp := grades[0]

	for i := 1; i < len(grades); i++ {
		if grades[i] >= temp {
			temp = grades[i]
		}
	}

	return temp
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}
	
	var average float64
	var min int
	var max int

	average = findAvg(s.Grades...)
	min = findMin(s.Grades...)
	max = findMax(s.Grades...)

	return average, min, max // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{},
	})
	fmt.Println(avg, min, max)
}
