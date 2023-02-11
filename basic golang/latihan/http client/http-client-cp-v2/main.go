package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://animechan.vercel.app/api/quotes/anime?title=naruto", nil)
    if err != nil {
        panic(err)
    }

	resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

	responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

	var a []Animechan

	err = json.Unmarshal(responseData, &a)
    if err != nil {
        return []Animechan{}, err
    }

	fmt.Println("get \n", responseData)

	// Hit API https://animechan.vercel.app/api/quotes/anime?title=naruto with method POST:
	return a, nil // TODO: replace this
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)

    // http.Post Implementation:
    resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
    
	//Handle Error
    if err != nil {
        log.Fatalf("An Error Occured %v", err)
    }
    defer resp.Body.Close()
    
	//Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
	
	var p Postman

	err = json.Unmarshal(body, &p)
    if err != nil {
        return p, err
    }

	// Hit API https://postman-echo.com/post with method POST:
	return p, nil // TODO: replace this
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
