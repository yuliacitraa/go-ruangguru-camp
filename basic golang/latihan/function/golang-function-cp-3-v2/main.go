package main

import (
	"fmt"
	"sort"
	"strings"
)


func FindShortestName(names string) string {
	var separator string
	for _, val := range names {
		char := fmt.Sprintf("%c", val)
		if char == " " {
			separator = " "
			break
		}else if char == ";" {
			separator = ";"
			break
		}else if char == "," {
			separator = ","
			break
		}
	}
	xy:=strings.Split(names, separator)
	var sliceName []string
	var sliceLen []int
	for _, name := range xy {
		sliceLen = append(sliceLen, len(name))
	}
	sort.Ints(sliceLen)
	for _, name := range xy {
		var x string

		if len(name) <= sliceLen[0] {
			x = name
			sliceName = append(sliceName, x)
		}
	}

	sort.Strings(sliceName)	
	return sliceName[0]
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
	fmt.Println(FindShortestName("Ari,Aru,Ara,Andi,Asik"))                // "Ara"
	fmt.Println(FindShortestName("Hanif Joko"))                			// "Ara"
}
