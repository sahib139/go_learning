package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web Post Request to server learning")
	performPostRequest()
}

func performPostRequest() {
	const url string = "http://localhost:8000/post"

	//fake json payload
	reqBody := strings.NewReader(`
		{
			"name":"sahib singh",
			"age":"22"
		}
	`)

	response, err := http.Post(url, "application/json", reqBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	resBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(resBody))
}
