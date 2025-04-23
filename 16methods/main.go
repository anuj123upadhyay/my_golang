package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritance in golang
	//no super or parent in go

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 22}
	fmt.Println("Hitesh details:", hitesh)
	fmt.Printf("Hitesh details: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v \n", hitesh.Name, hitesh.Email )
	hitesh.getStatus()
	hitesh.NewMail()
	fmt.Printf("Name is %v and email is %v \n", hitesh.Name, hitesh.Email )
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int

}


func (u User) getStatus(){
	fmt.Println("User status is", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is", u.Email)
}
