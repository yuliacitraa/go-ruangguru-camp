package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	tld := strings.Split(email, ".")
	domain := strings.Split(tld[0], "@")


	if len(tld) > 2 {
		return	"Domain: "+domain[1]+" dan TLD: " + tld[1] +"."+tld[2]

	}

	return	"Domain: "+domain[1]+" dan TLD: " + tld[1]
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.com"))
}
