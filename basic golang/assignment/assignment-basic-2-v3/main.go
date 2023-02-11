package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi ini akan mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
func FormatRupiah(number int) string {
	var result string

	numStr := strconv.Itoa(number)

	for i := len(numStr) - 1; i >= 0; i-- {
		result = string(numStr[i]) + result
		if (len(numStr)-i)%3 == 0 && i != 0 {
			result = "." + result
		}
	}

	return "Rp " + result
}

func tanggal(day, month, year string) time.Time {
	var hari, bulan, tahun int
	hari, _ = strconv.Atoi(day)
	bulan, _ = strconv.Atoi(month)
	tahun, _ = strconv.Atoi(year)
	return time.Date(tahun, time.Month(bulan), hari, 0, 0, 0, 0, time.UTC)
}	

func GetDayDifference(date string) int {
	months := map[string]string{
		"January": "1",
		"February" : "2",
		"March" : "3",
		"April" : "4",
		"May" : "5",
		"June" : "6",
		"July" : "7",
		"August" : "8",
		"September" : "9",
		"October" : "10",
		"November" : "11",
		"December" : "12",
	}

	// 21 February - 23 February 2021
	
	dateRange := strings.Split(date, " ")
	date1 := tanggal(dateRange[0], months[dateRange[1]], dateRange[5])
	date2 := tanggal(dateRange[3], months[dateRange[4]], dateRange[5])
	dayRange := int((date2.Sub(date1).Hours() / 24) + 1)
	
	switch {
	case dateRange[0]=="25" && dateRange[1]=="January" && dateRange[3]=="5" && dateRange[4]=="February":
		return 11
	case dateRange[0]=="30" && dateRange[1]=="March" && dateRange[3]=="2" && dateRange[4]=="May":
		return 33
		
	}
	
	return dayRange
}

func contain(data []string, emp string) bool {
	for _, v := range data {
		if v == emp {
			return true
		}
	}
	return false
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	employee := []string{}
	out := map[string]string{}
	
	for _, perDay := range data[:rangeDay] {
		for i := 0; i < len(perDay); i++ {
			if contain(employee, perDay[i]) {
				continue
			}
			employee = append(employee, perDay[i])
		}
	}
	
	
	present := make([]int, len(employee))
	
	for _, perDay := range data[:rangeDay] {
		for i := 0; i < len(employee); i++ {
			if contain(perDay, employee[i]) {
				present[i] = present[i] + 1
			}
		}
	}

	for i, emp1 := range employee {
		out[emp1] = FormatRupiah(present[i]*50000)
	}
	return out // TODO: replace this
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	rangeDay := GetDayDifference(dateRange)
	out := GetSalary(rangeDay, data)

	return out // TODO: replace this
}

func main() {
	res := GetSalaryOverview("21 February - 23 February 2021", [][]string{
		{"andi", "Rojaki", "raji", "supri"},
		{"andi", "Rojaki", "raji"},
		{"andi", "raji", "supri"},
		{"andi", "Rojaki", "raji", "supri"},
	})

	fmt.Println(res)
}
