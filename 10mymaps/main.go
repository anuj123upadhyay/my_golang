package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps in golang")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"
	languages["go"] = "Golang"

	fmt.Println("List of all Languages: ", languages)
	fmt.Println("JS shorts for: ", languages["js"])


	//delete some values
	delete(languages , "rb")
	fmt.Println("List of all Languages: ", languages)


	//loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For Keys: %v , value is %v\n", key, value )
	}
   
	// with comma okk syntax
	for _, value := range languages {
		fmt.Printf("For Keys: v , value is %v\n", value )
	}


}
