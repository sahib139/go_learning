package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Web Post Request to server learning")
	performPostRequest()
}

func performPostRequest() {
	const URL string = "http://localhost:8000/postform"

	//form data
	data := url.Values{}

	data.Add("name", "sahib singh")
	data.Add("Age", "22")

	response, err := http.PostForm(URL, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	resBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(resBody))
}
