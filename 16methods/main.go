package main

import "fmt"

func main() {
	fmt.Println("Learning Methods")

	user1 := User{Name: "sahib", Email: "ssok@12.com", Age: 20}
	fmt.Println(user1)
	result := user1.voteIllegible()
	fmt.Println(result)
	user1.changeEmail("ok@12@23.com")
	fmt.Println(user1)
	user1.updateEmail("ok@12@23.com")
	fmt.Println(user1)
}

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) voteIllegible() bool {
	if u.Age > 10 {
		return true
	} else {
		return false
	}
}

func (u User) changeEmail(newEmail string) {
	u.Email = newEmail 
	// this will be ineffective since it passes the copy of user , and in that we change
	// email value and hence it don't reflect in the original user object.
}

func (u *User) updateEmail(newEmail string) {
	(*u).Email = newEmail // by passing the pointer we able do the changes
}
