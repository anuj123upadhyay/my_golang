package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println("Days are", days)


	// for i:=0;i< len(days); i++ {
	// 	fmt.Println(days[i])
	// }


	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	// for _, day := range days {
	// 	fmt.Printf("index is  and value is %v\n", day)
	// }



	// syntax somewhat similar to while
	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		} 

		// if rougueValue == 5 {
		// 	break
		// }

		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is :", rougueValue)
		rougueValue++


}

//goto statement
	lco:
		fmt.Println("Jumping to LearnCodeOnline.in")

}