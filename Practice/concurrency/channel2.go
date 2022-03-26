package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channel in golang")

	mych := make(chan int, 2) // buffer channel with 2
	wg := &sync.WaitGroup{}
	// fmt.Println(<-mych)
	// mych <- 5
	wg.Add(2)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mych <- 5
		mych <- 6
		close(ch)
		wg.Done()
	}(mych, wg)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// need buffer channel
		// fmt.Println(<-mych)
		//var, isChanelOpen := <-mych
		fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)

	wg.Wait()
}
