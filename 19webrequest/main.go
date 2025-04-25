package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://rohitupadhyay.me"

func main() {
	fmt.Println("Welcome to Web requestin golang")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of Type %T\n", response)

	defer response.Body.Close() // callers's responsibiity to close the connections

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println("Content of the page is: ", content)

}
