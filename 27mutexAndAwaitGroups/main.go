package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition ---------")
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}



	var score = []int{0}


	// wg.Add(1)
	wg.Add(4)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One Goroutine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two Goroutine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three Goroutine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("four Goroutine")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println("Score:", score)
}
