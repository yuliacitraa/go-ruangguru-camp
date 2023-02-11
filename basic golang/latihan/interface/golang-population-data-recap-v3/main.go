package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{}  {
	if len(data) == 0 {
		return []map[string]interface{}{}
	}

	var list []string
	mapping := make(map[string]interface{})
	result := []map[string]interface{}{}
	for _, v := range data {
		list = strings.Split(v, ";")
		mapping["name"] = list[0]
		mapping["age"], _ = strconv.Atoi(list[1]) 
		mapping["address"] = list[2]

		if list[3] != ""{
			mapping["height"], _ = strconv.ParseFloat(list[3], 64)
		}
		if list[4] != ""{
			mapping["isMarried"], _ = strconv.ParseBool(list[4])
		}
		result = append(result, mapping)
		mapping = make(map[string]interface{})
	}
	return result
}

func main()  {
	data := []string{"Budi;23;Jakarta;;true","Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	fmt.Println(PopulationData(data))
}