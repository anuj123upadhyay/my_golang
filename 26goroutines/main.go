package main
// go keyword is used to declare a goroutine => go keyword

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //usually these are pointers

var mut sync.Mutex //mutual exclusion lock
  // usually these are pointers
func main() {
	// fmt.Println("Welcome to Goroutines")
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://rohitupadhyay.me",
		"https://google.com",
		"https://facebook.com",
		"https://x.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)

	}
	wg.Wait()
	fmt.Println("Signals:", signals)
}

// func greeter(s string) {

// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}

// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	} else {
		mut.Lock()
		signals = append(signals,endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
	}

}
