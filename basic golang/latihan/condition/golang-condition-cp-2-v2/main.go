package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	tinggi := float64(height)
	selisih := tinggi-100
	if gender == "laki-laki" {
		bmi := (selisih-(selisih*0.1))
		return bmi
	}else {
		bmi := (selisih-(selisih*0.15))
		return bmi
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 170))
	fmt.Println(BMICalculator("perempuan", 165))
}
