package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of go lamguage")
	presentTime := time.Now()
	presentTimeNano := time.Now().Nanosecond()
	fmt.Println(presentTime)
	fmt.Println(presentTimeNano)
	fmt.Println(presentTime)

	fmt.Println( presentTime.Format("01-03-2006 15:03:23 Monday"))


	createDate := time.Date(2020,time.August , 11, 21, 21 , 0, 0, time.UTC)
	fmt.Println("Created date is ", createDate)
	fmt.Println("Created date is ", createDate.Format("01-03-2006 15:03:23 Monday"))


}
