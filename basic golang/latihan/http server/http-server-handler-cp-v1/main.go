package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	t := time.Now()
	
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
	}

	
	return handler // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
