package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://rohitupadhyay.me:3000/learn?coursename=reactjs&paymentid=dgh123gvdg"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println("myurl => ", myurl)

	// parsing the url => extracting the info from url
	result, _ := url.Parse(myurl)

	fmt.Println("result : ", result)
	fmt.Println("result scheme is : ", result.Scheme)
	fmt.Println("result Host is : ", result.Host)
	fmt.Println("result HostName is : ", result.Hostname())
	fmt.Println("result Path is : ", result.Path)
	fmt.Println("result PORT is : ", result.Port())
	fmt.Println("result Raw Query is : ", result.RawQuery)//=> is like the parameter passed in the url
	fmt.Println("result ForceQuery is : ", result.ForceQuery)

	qparams := result.Query()
	fmt.Printf("the type of Query params is : %T\n", qparams)
	fmt.Println("qparams is : ", qparams)

	fmt.Println("Coursename is :", qparams["coursename"])
	fmt.Println("Coursename is :", qparams["paymentid"])

	for _, value := range qparams {
		fmt.Println("Value is : ", value)
	}

	partsOfUrl :=&url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:    "/tutcss",
		RawQuery: "user=rohit",
		
	}
	anotherURL := partsOfUrl.String()
	fmt.Println("Another URL is : ", anotherURL)

}
