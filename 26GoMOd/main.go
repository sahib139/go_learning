package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Go Mod learning!")
	goHttpSever()
}

func goHttpSever() {
	r := mux.NewRouter()
	r.HandleFunc("/", routerFunc)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func routerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>My name is sahib</h1>"))
}
