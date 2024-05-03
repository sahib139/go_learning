package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	College   string   `json:"college"`
	CollegeId string   `json:"-"`
	Courses   []string `json:"courses,omitempty"`
}

func main() {
	fmt.Println("Json Data creation!!")

	jsonDataCreation()
}

func jsonDataCreation() {

	users := []user{
		{Name: "Sahib Singh", Age: 11, College: "cool", CollegeId: "coolSahib", Courses: []string{"DBMS", "DSA", "Compiler Design", "OS"}},
		{Name: "ok abc", Age: 22, College: "awesome", CollegeId: "awesomeOk", Courses: []string{"DBMS", "DSA"}},
		{Name: "xyz", Age: 33, College: "amazing", CollegeId: "amazingXyz", Courses: nil},
	}

	finalJSON, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}
