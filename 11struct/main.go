package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	user1 := User{Name: "Sahib Singh", Email: "sahib.2125cse1145@kiet.edu", Age: 22}
	fmt.Printf("%v\n %+v \n Name: %v", user1, user1, user1.Name)
}
