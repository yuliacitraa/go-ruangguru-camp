package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
	
    scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
		lines = append(lines, scanner.Text())
    }

	if len(lines) == 0 {
		return []string{}, nil
	}

    return lines, scanner.Err()
}

func CalculateProfitLoss(data []string) string {
	var total, income, expense int
	var lastDate string
	for _, v := range data {
		split := strings.Split(v, ";")
		amount, err := strconv.Atoi(split[2])
		if err != nil {
			panic(err)
		}
		
		if split[1] == "income" {
			income += amount
		}
		if split [1] == "expense" {
			expense += amount 
		}

		lastDate = split[0]
		
	}
	total = income - expense

	if income > expense {
		return fmt.Sprintf("%s;profit;%d", lastDate, total)
	}

	if income < expense {
		return fmt.Sprintf("%s;loss;%d", lastDate, total * (-1))
	}
	
	return ""
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
