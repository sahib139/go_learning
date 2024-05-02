package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web Get Request learning!")
	performGetRequest()
}

func performGetRequest() {

	const url string = "http://localhost:8000/get"
	request, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer request.Body.Close()

	content, _ := io.ReadAll(request.Body)

	fmt.Println(string(content), request.ContentLength)

	// another way to read string
	var requestBody strings.Builder
	byteCount, _ := requestBody.Write(content)

	fmt.Printf("The byteCount of request body is %v\n", byteCount)
	fmt.Println(requestBody.String())

}
