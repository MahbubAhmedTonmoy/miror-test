package main

import (
	"fmt"
	"sync"
)

//go run --race .
//exit status 66

func main() {
	fmt.Println("race condition")

	waitGroup := &sync.WaitGroup{}
	mute := &sync.Mutex{}
	var score = []int{0}

	waitGroup.Add(3)
	//annonymous function
	go func(waitGroup *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("1st R")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		waitGroup.Done()
	}(waitGroup, mute)

	go func(waitGroup *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("2nd R")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		waitGroup.Done()
	}(waitGroup, mute)

	go func(waitGroup *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("3rd R")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		waitGroup.Done()
	}(waitGroup, mute)

	waitGroup.Wait()
	fmt.Println(score)
}
