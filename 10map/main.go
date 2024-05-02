package main

import "fmt"

func main() {
	mp := make(map[string]string)
	mp["FirstName"] = "sahib"
	mp["SecondName"] = "singh"
	fmt.Println(mp)
	fmt.Printf("FullName is :%s %s\n", mp["FirstName"], mp["SecondName"])
	// delete(mp, "SecondName")
	// fmt.Println(mp)

	for key, val := range mp {
		fmt.Printf("The key is :%s and the value is %s\n", key, val)
	}

}
