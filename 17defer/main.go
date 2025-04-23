package main

import "fmt"

func main() {
	defer fmt.Println("World")// whatever written with the defer , will be executed at the last of the function
	//or we can say that this line gets cut here and paste it just before the '}' of the function
	//they aslo follow LIFO order ( Last in First out)
	defer fmt.Println("One ")
	defer fmt.Println("Two ")
	//first defer statement will be at the end to be executed
	fmt.Println("Hello ")
	myDefer()
}

func myDefer() {
	for i:=0;i<5;i++{
		defer fmt.Println(i) // this will print all the values in reverse order
	}
}
