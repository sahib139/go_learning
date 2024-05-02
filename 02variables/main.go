package main

import (
	"fmt"
)

var Uoo = "sahib" // if start with capital letter then it is public variable

func main() {
	var name string = "sahib"
	fmt.Println(name)
	fmt.Printf("THe type of variable is %T\n", name)

	var smallUnsignedInt uint8 = 255
	fmt.Println(smallUnsignedInt)
	fmt.Printf("THe type of variable is %T\n", smallUnsignedInt)

	var smallSignedInt int8 = 120
	fmt.Println(smallSignedInt)
	fmt.Printf("THe type of variable is %T\n", smallSignedInt)

	var smallFloat float32 = 255.0032
	fmt.Println(smallFloat)
	fmt.Printf("THe type of variable is %T\n", smallFloat)

	var booleanValue bool
	fmt.Println(booleanValue)
	booleanValue = true
	fmt.Printf("THe type of variable is %T\n", booleanValue)

	var number = 10
	fmt.Println(number)
	fmt.Printf("THe type of variable is %T\n", number)

	newNumber := 1.141
	fmt.Println(newNumber)
	fmt.Printf("THe type of variable is %T\n", newNumber)

	fmt.Println(Uoo)
}
