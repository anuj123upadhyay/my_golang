package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channesl in golang------")
	// channels are used to communicate between goroutines

	myCh := make(chan int, 2) // create a channel of type int

	wg := &sync.WaitGroup{}

	// myCh <- 5 // send data to the channel
	// fmt.Println(<-myCh)
	wg.Add(2)
	//Receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh // receive data from the channel
		fmt.Println("Value:", val, "isChannelOpen:", isChannelOpen)
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	//send only channel
	go func(ch chan <-int, wg *sync.WaitGroup) {

		myCh <- 5 // send data to the channel
		myCh <- 10
		close(myCh) // close the channel
		wg.Done()
	}(myCh, wg)
	wg.Wait()

}
