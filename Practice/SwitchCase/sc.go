package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case in go lang")
	rand.Seed(time.Now().Unix())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 space")
	case 3:
		fmt.Println("you can move 3 space")
	case 4:
		fmt.Println("you can move 4 space")
		fallthrough
	case 5:
		fmt.Println("you can move 5 space")
	case 6:
		fmt.Println("you can move 6 space and roll again")
	default:
		fmt.Println("out")
	}

}
