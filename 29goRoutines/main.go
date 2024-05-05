package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var statusCodes []int

func main() {

	websiteList := []string{
		"https://instagram.com",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, val := range websiteList {
		wg.Add(1)
		go getStatusCode(val)
	}
	wg.Wait()
	fmt.Println(statusCodes)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	response, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status code of endpoint:%v is %v\n", endpoint, response.StatusCode)

	mut.Lock()
	statusCodes = append(statusCodes, response.StatusCode)
	mut.Unlock()
}
