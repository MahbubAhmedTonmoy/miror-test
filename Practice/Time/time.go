package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("welcome to time study")

	prestTime := time.Now()
	fmt.Println(prestTime)
	fmt.Println(prestTime.Date())
	fmt.Println(prestTime.Format("2006/01/02 15:04:05 Monday"))

	createTime := time.Date(2022, time.August, 15, 12, 10, 10, 0, time.UTC)
	fmt.Println(createTime)

	fmt.Println(time.Unix(0, 0))
	fmt.Println(os.Args[0])
	fmt.Println(os.Args)
}
