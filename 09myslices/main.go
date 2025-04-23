package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in golang")

	var fruitList = []string{"Apple", "Mango", "Peach"}

	fmt.Printf("Type of Fruitlist is %T\n", fruitList)


	fruitList = append(fruitList, "Orange","Banana")
	fmt.Println("fruitList:", fruitList)



	// fruitList = append(fruitList[1:3])
	// fmt.Println("fruitList:", fruitList)
	fruitList = append(fruitList[0:3])
	fmt.Println("fruitList:", fruitList)

	
	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 951
	highScore[2] = 434
	highScore[3] = 334
	// highScore[4] = 34

	highScore = append(highScore, 555, 666, 777)


	fmt.Println("High Scores:", highScore)


	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println("highScores:", highScore)


	//how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("courses", courses)



}
