package main

import (
	"bufio" 
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("welcome to mahbub online")
	fmt.Println("input 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter")

	input, _ := reader.ReadString('\n')

	fmt.Printf(input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating + 1)
	}

}
