package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var waitGroup sync.WaitGroup //pointer

var mute sync.Mutex //pointer

func main() {
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
		"https://gitlab.com",
		"https://bitbucket.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		waitGroup.Add(1)
	}
	waitGroup.Wait()
	fmt.Println(signals)
}

func getStatusCode(endPoint string) {

	defer waitGroup.Done()
	res, err := http.Get(endPoint)
	if err != nil {
		fmt.Println("error occure")
	} else {
		mute.Lock()
		signals = append(signals, endPoint)
		mute.Unlock()
		fmt.Printf("endpoint is, %s ---> %d\n", endPoint, res.StatusCode)
	}
}
