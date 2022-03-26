package channel

import (
	"fmt"
	"time"
)

//one direction (ch <-chan int)
//both direction
func receiver(ch chan int) {
	// for {
	// 	i, ok := <-ch
	// 	if ok == false {
	// 		fmt.Println(i, ok, "<-- loop break")
	// 		break
	// 	} else {
	// 		fmt.Println(i, ok)
	// 	}
	// }

	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println("channel close")
	ch <- 100 // send from reciver
}
func main() {
	fmt.Println("Channels")

	ch := make(chan int) //  create channel with capacity  ch:= make(chan int ,2)
	go receiver(ch)
	ch <- 42 // write in channel
	ch <- 43
	ch <- 44
	ch <- 45
	fmt.Println(<-ch)
	close(ch)
	time.Sleep(1000 * time.Microsecond)
	//<-ch // read from channel
}
