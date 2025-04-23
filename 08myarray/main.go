package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array in golang")

	var fruitList [6]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Mango"
	fruitList[3] = "Orange"
	fruitList[5] = "Grapes"
	fmt.Println("Fruit list is ", fruitList)
	fmt.Println("Fruit list is ", len(fruitList))


	var VegList = [3]string{"Potato","Tomato", "Onion"}
	fmt.Println("Veg list is ", VegList)
	fmt.Println("Veg list is ", len(VegList))
}
