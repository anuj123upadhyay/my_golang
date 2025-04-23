package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Swict and Case in GO Lang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Number", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You rolled a 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
		fallthrough //  this will execute the next case as well

	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and roll dice again")
	default:
		fmt.Println("What was that")
	}

}
