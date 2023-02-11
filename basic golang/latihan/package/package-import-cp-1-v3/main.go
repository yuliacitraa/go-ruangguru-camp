package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

// func AdvanceCalculator(calculate string) float32 {
// 	Calc := internal.NewCalculator(0)
// 	input := strings.Split(calculate, " ")
// 	var number, angkaPertama float32
// 	var symb func(float32) = nil
// 	var isAngka = false
// 	for _, v := range input {
// 		switch v{
// 		case "+":
// 			symb = Calc.Add
// 		case "-":
// 			symb = Calc.Subtract
// 		case "*":
// 			symb = Calc.Multiply
// 		case "/" :
// 			symb = Calc.Divide
// 		default :
// 			isAngka = true
// 		}

// 		if isAngka {
// 			n, _ := strconv.ParseFloat(v, 32)
// 			if symb == nil {
// 				angkaPertama = float32(n)
// 				Calc = internal.NewCalculator(angkaPertama)
// 			} else {
// 				number = float32(n)
// 			}
// 		}
// 		fmt.Println(number, Calc.Result())
// 		if (v != "+" || v != "-" || v != "*" || v != "/") && symb != nil {
// 			symb(number)
// 		}
// 	}
// 	return Calc.Result()
// }

func AdvanceCalculator(calculate string) float32  {
	split := strings.Split(calculate, " ")
	s, _ := strconv.Atoi(split[0])
	calc := internal.NewCalculator(float32(s))
	new := split[1:]

	for i, v := range new {
		if i % 2 == 0 {
			num, _ := strconv.Atoi(new[i+1])
			switch v {
			case "+":
				calc.Add(float32(num))
			case "-":
				calc.Subtract(float32(num))
			case "*":
				calc.Multiply(float32(num))
			case "/":
				calc.Divide(float32(num))
			}
		}
	}

	return calc.Result()

}

func main() {
	// res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")
	res := AdvanceCalculator("10 + 2 + 3 + 4 * 2")


	fmt.Println(res)
}
