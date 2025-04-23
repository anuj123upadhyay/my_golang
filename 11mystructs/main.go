package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritance in golang
	//no super or parent in go

	hitesh := User{"Hitesh","hitesh@go.dev",true, 22}
	fmt.Println("Hitesh details:", hitesh)
	fmt.Printf("Hitesh details: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}


