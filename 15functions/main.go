package main

import "fmt"

func main() { // main function is the entry point of the program
	fmt.Println("Welcome to Functions in Go Lang")
	greeter()
	// func greeterTwo(){
	// 	fmt.Println("Another method")
	// }
	// greeterTwo()    // function inside the functions are not allowed in go
	greeterTwo()

	result := adder(3, 5)

	fmt.Println("Result is : ", result)

	proResult , myMessage:= proAdder(2,5,8,7)
	fmt.Println("Pro Result is : ", proResult)
	fmt.Println("Pro Message is : ", myMessage)

}

func adder(val1 int, val2 int) int {

	return val1 + val2
}

func proAdder(values ...int)(int, string){
	total :=0
	for _, val := range values{
		total += val
	}
	return total , "Hi from Pro result"
}



//==========================
func greeter() {
	fmt.Println("Namastey from Golang")
}

func greeterTwo() {
	fmt.Println("Another method")
}

// anonymous function
// func (){
// 	fmt.Println("I am anonymous function")
// }()
