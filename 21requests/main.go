package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Welcome to web verb video- LCO")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "https://rohitupadhyay.me/"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code ", response.StatusCode)

	fmt.Println("Content Length ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("Response ", string(content))

	byteCount, _ := responseString.Write(content) // another way of doing the same thing
	fmt.Println("byteCount ", byteCount)
	fmt.Println("Response in string ", responseString.String())

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"learncodeonline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json:", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/portform"
	// fake formdata

	data := url.Values{}
	data.Add("firstname", "Rohit")
	data.Add("lastname", "Upadhyay")
	data.Add("email", "	rohit@go.dev")

	response , err := http.PostForm(myurl, data)
	if err!= nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	fmt.Println("Status Code ", response.StatusCode)	
}
