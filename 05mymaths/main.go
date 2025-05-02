package main

import (
	"fmt"
	"math/big"
	// "math/rand"

	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to Math sin Golang")

	// var mynumberone int =2
	// var mynumbertwo float64 =4

	// fmt.Println("the sum is :", mynumberone

	//random number generation => using math/rand package
	//it panic when the number is less than 0
	// rand.Seed(30)
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("rand.Intn(5) => ", rand.Intn(5)) // 0 to 4

	//random number generation  => using crypto/rand package
	//it use cryptographic algo for generation

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("My random number is : ", myRandomNum) // 0 to 4

}
