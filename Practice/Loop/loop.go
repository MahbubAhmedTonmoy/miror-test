package main

import "fmt"

func main() {
	fmt.Println("loop in go")
	days := []string{"sunday", "monday", "tue", "wen", "thus"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for _, i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index, day:=range days{
	// 	fmt.Printf("index is %v and value is %v \n",index, day)
	// }

	value := 1

	for value < 10 {
		if value == 2 {
			goto jump
		}
		if value > 5 {
			value++
			continue //break
		}
		fmt.Println("value is :", value)
		value++
	}
jump:
	fmt.Printf("jum from loop")
}
