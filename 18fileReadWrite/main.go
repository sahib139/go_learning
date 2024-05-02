package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("File Read-Write learning")

	file, err := os.Create("MyTextFile.txt")

	checkNilError(err)

	length, err := file.WriteString("hey my name is sahib")
	// length, err := io.WriteString(file, "hey my name is sahib")

	checkNilError(err)
	fmt.Println(length)

	defer file.Close()
	readFile("./MyTextFile.txt")
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)

	checkNilError(err)

	fmt.Println(string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
