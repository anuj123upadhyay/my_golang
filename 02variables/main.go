package main

import "fmt"


// jwtToken = 3000000  // outside the function you are not allowed to use walrus opeartor

// var jwt int =30000  // this is allowed


const LoginToken  string = "anujupadhyay"  // this is the practice to decalre a Public Constant the first letter should always be the "Capital letter"
func main() {
	var username string = "Anuj Kumar Upadhyay"
	var age int  = 21
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n",username)
	fmt.Printf("Variable is of type : %T \n",age)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n",isLoggedIn)


	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n",smallVal)


	var smallFloat float32 = 255.455445123 
	var bigFloat float64 = 255.455445123 
	fmt.Println(smallFloat)
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type : %T \n",smallFloat)
	fmt.Printf("Variable is of type : %T \n",bigFloat)



	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n",anotherVariable)

	var emptyString string
	fmt.Println(emptyString)
	fmt.Printf("Variable is of type : %T \n",emptyString)


	//implicit types

	var website = "anujupadhyay.me"
	fmt.Printf("Variable is of type : %T \n",website)

	//no var style

	numberOfUser := 300000.0 // := using "walrus operator"
	fmt.Println("Number of users is ", numberOfUser)



	fmt.Println("Login Token is ", LoginToken)
	fmt.Printf("Variable is of type : %T \n",LoginToken)
	
}