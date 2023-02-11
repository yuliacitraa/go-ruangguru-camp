package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	out := make(map[string][]string)
	for i := 0; i < len(data); i++ {
		split := strings.Split(data[i], "-")
		key := split[0]
		num, _ := strconv.Atoi(split[1])
		desc := split[2]
		value := split[3]
	
		var ok bool
		_, ok = out[key]
		if !ok {
			out[key] = make([]string, 0)
		}
	
		switch  {
		case desc == "first" :
			if num >= len(out[key]) {
				out[key] = append(out[key], "")
			}
			out[key][num] = value + out[key][num]
		case desc == "last" :
			if num >= len(out[key]) {
				out[key] = append(out[key], "")
			}
			out[key][num] = out[key][num] + " " + value
		}
	}	
	return out // TODO: replace this
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
