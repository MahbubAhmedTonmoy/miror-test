package concurrency

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main thread")
	//go hello(1)
	//go hello(2)

	//annonymous
	msg := "hello"
	go func() {
		time.Sleep(10 * time.Microsecond)
		fmt.Println(msg)
	}()
	msg = "world"
	//

	fmt.Println("main thread end")
	time.Sleep(100 * time.Microsecond)
}

func hello(gen int) {
	for i := 0; i < 5; i++ {
		fmt.Println(gen, ": Hello")
	}
}
