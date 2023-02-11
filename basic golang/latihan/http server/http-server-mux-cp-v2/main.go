package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		writer := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
		w.Write([]byte(writer))
	} // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		name := r.URL.Query().Get("name")
		if name == "" {
			writer := fmt.Sprint("Hello there")
			w.Write([]byte(writer))
		} else {

			writer := fmt.Sprintf("Hello, %s!", name)
			w.Write([]byte(writer))
		}
	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
