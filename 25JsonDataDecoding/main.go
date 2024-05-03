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
	fmt.Println("JSON data decoding !")
	JSONdecoding()
}

func JSONdecoding() {
	var users user

	jsonString := []byte(`
		{
			"name":"Sahib Singh",
			"age":20,
			"college":"ABC",
			"courses":[
				"DSA",
				"OS"
			]
		}
	`)

	checkValidation := json.Valid(jsonString)

	if checkValidation {
		json.Unmarshal(jsonString, &users)
		fmt.Printf("%#v\n", users)
		fmt.Printf("%v,%v,%v,%v\n", users.Name, users.Age, users.College, users.Courses)
	} else {
		fmt.Println("The json data is not valid")
	}

	var usersData map[string]interface{}
	if checkValidation {
		json.Unmarshal(jsonString, &usersData)
		fmt.Printf("%#v\n", usersData)
		for key, val := range usersData {
			fmt.Printf("key :%v,val :%v,key_type:%T,val_Type:%T\n", key, val, key, val)
		}
		fmt.Println(usersData["name"], usersData["courses"])
	} else {
		fmt.Println("The json data is not valid")
	}

}
