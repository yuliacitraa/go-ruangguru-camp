package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	out := make([][]string, 0, len(mapData))
	for key, value := range mapData {
		out = append(out, []string{key, value})
	}

	return out // TODO: replace this
}

func main()  {
	// input:= []string{"hello": "world", "John": "Doe" , "age":"14"}
	fmt.Println(MapToSlice(map[string]string{"hello": "world", "John": "Doe" , "age":"14"}))
}