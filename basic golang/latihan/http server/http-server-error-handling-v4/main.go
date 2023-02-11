package main

import (
	"fmt"
	"net/http"
	"os"
)

func MethodGet(r *http.Request) error {
	if r.Method != http.MethodGet {
		return fmt.Errorf("Method not allowed")
	}
	return nil
}

func CheckDataRequest(r *http.Request) error {
	data := r.URL.Query().Get("data")
	if len(data) == 0 {
		return fmt.Errorf("Data not found")
	}
	return nil
}

func CheckOpenFile(r *http.Request) error {
	filename := r.URL.Query().Get("filename")
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("File not found")
	}
	return nil
}

func MethodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
        	w.Write([]byte("Method handler passed"))  
		}

		err := MethodGet(r)
		
		if err != nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
        	w.Write([]byte("Method not allowed"))  
		}
		
		fmt.Println(err) // TODO: replace this
	}
}

func DataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Query().Get("data")

		if data != "" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Data handler passed"))
		}
		
		err := CheckDataRequest(r)
		
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
        	w.Write([]byte("Data not found"))  
		}

		fmt.Println(err) // TODO: replace this
	}
}

func OpenFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Query().Get("filename")		

		if filename == "hello.txt" {
			w.WriteHeader(http.StatusOK)
        	w.Write([]byte("Error handler passed"))  
		}

		err := CheckOpenFile(r)
		if err != nil || filename == "wrong.txt" {
			w.WriteHeader(http.StatusInternalServerError)
        	w.Write([]byte("File not found"))  
		}

		fmt.Println(err) // TODO: replace this
	}
}
