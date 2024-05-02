package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "http://13.60.19.44:3001/html/signIn.html"

func main() {
	response, err := http.Get(URL)

	checkNilError(err)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	checkNilError(err)

	fmt.Println(string(body))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
