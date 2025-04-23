package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files in golang")

	content := "This needs to go in file - LearnCodeOnline.in"
	file, err := os.Create("./myfile.txt") // this will create a file with the name myfile.txt

	if err != nil {
		panic(err) // this will stop the execution of the programans show off the error
	}

	length, err := io.WriteString(file, content) // this will return the length

	checkNilError(err)

	fmt.Println("length is : ", length)
	defer file.Close() //recommended to use defer to close the file

	readFile("./myfile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	checkNilError(err)
	fmt.Println("Text data inside the file is: ", dataByte)
	fmt.Println("Text data inside the file is: ", string(dataByte)) // this will convert the byte array to string
}

func checkNilError(err error) {
	if err != nil {
		panic(err) // this will stop the execution of the programans show off the error
	}
}
