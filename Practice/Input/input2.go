package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("welcome to map input in go lang")
	//map
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for j, n := range counts {
		if n > 1 {
			fmt.Println(n, j)
		}
	}
}
