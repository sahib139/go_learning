package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/sahib139/go_Basic_API_using_mongo/routers"
)

var wg sync.WaitGroup

func main() {
	router := routers.Router()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(http.ListenAndServe(":4000", router))
	}()

	fmt.Println("Server Started at Port: 4000")

	wg.Wait()
}
