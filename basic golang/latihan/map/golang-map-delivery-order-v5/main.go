package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {
	out := make(map[string]float32)

	
	cityPerDay := map[string][]string {
		"JKT" : {"senin", "selasa", "rabu", "kamis", "jumat","sabtu"},
		"BDG" : {"rabu", "kamis", "sabtu"},
		"BKS" : {"selasa", "kamis", "jumat"},
		"DPK" : {"senin", "selasa"},
	}
	
	adminPerDay := map[string]float32 {
		"senin" : 0.1,
		"selasa" : 0.05,
		"rabu" : 0.1,
		"kamis" : 0.05,
		"jumat" : 0.1,
		"sabtu" : 0.05,
	}
	
	for i := 0; i < len(data); i++ {
		split := strings.Split(data[i], ":")
		firstName := split[0]
		lastName:= split[1]
		p := split[2]
		pay, _ := strconv.ParseFloat(p, 32)
		pays := float32(pay)
		city := split[3]
		

		passed := false
		for _, d := range cityPerDay[city] {
			if d == day {
				passed = true
			}
		}

		admin := pays + (pays * adminPerDay[day])

		if passed {
			out[firstName+"-"+lastName] = admin
		}
	}
	
	return out // TODO: replace this
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
