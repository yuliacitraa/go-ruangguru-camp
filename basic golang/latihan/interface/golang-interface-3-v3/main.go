package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func formatDate(hour, minute int) string{
	var out string
	if hour < 12 {
		out = fmt.Sprintf("%02d:%02d AM", hour, minute)
	} else if hour == 24{
		out = fmt.Sprintf("%02d:%02d AM", hour-12, minute)
	} else if hour > 12 {
		out = fmt.Sprintf("%02d:%02d PM", hour-12, minute)
	} else if hour == 12{
		out = fmt.Sprintf("%02d:%02d PM", hour, minute)
	}
	return out
}


func ChangeToStandartTime(time interface{}) string {
	var out string
	fmt.Println(time)

	switch time.(type) {
	case string:
		split := strings.Split(time.(string), ":")
		
		if len(split) !=2 || split[0]=="" || split[1]=="" {
			out = "Invalid input"
			break
		}
		
		hourStr := split[0]
		minStr := split[1]
		
		hour, _ := strconv.Atoi(hourStr)
		min, _ := strconv.Atoi(minStr)
		

		out = formatDate(hour, min)


	case []int :
		var timeSlice = time.([]int)
		if len(timeSlice) != 2 {
			out = "Invalid input"
			break
		}
		hour := timeSlice[0]
		min := timeSlice[1]

		out = formatDate(hour, min)


	case map[string]int :
		var timeMap = time.(map[string]int)
		hour := timeMap["hour"]
		min := timeMap["minute"]
		if len(timeMap) !=2 {
			out = "Invalid input"
			break
		}
		if _, ok := timeMap["second"]; ok {
			out = "Invalid input"
			break
		}
		out = formatDate(hour, min)


	case Time :
		hour := time.(Time).Hour
		min := time.(Time).Minute
		out = formatDate(hour, min)


	}
	return out // TODO: replace this
}

func main() {
	// fmt.Println(ChangeToStandartTime("2300"))
	// fmt.Println(ChangeToStandartTime("23:"))
	// fmt.Println(ChangeToStandartTime(Time{16, 0}))
	// fmt.Println(ChangeToStandartTime("12:00"))
	// fmt.Println(ChangeToStandartTime("14:00"))

	// fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime([]int{16}))
	// fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16}))

}
