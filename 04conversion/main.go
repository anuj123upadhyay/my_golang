package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza rating app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	// numRating, err := strconv.ParseFloat(input, 64)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)


	if err != nil {
		fmt.Println("Error while parsing the input, ", err)
	}else{
		fmt.Println("Added 1 to your rating, ", numRating + 1)
	}


}
