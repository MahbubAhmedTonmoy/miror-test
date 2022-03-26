package main

import "fmt"

func main() {
	fmt.Println("welcome")
	print()

	result := add(3, 4)
	fmt.Println(result)

	r2, message := proAdder(1, 2, 3, 4, 5)
	fmt.Println(r2, message)
	//annonymous function
	func() {
		fmt.Println("hello 2")
	}()

	//call function which receive annonymous function as argument
	value := func(p, q string) string {
		return p + q
	}
	PassFunction(value)

	//call return an anonymous function from another function
	rf := returnAnoFunc()
	fmt.Println(rf("mahbub", "Ahmed"))
}

//return an anonymous function from another function
func returnAnoFunc() func(i, j string) string {
	f := func(i, j string) string {
		return i + " " + j
	}
	return f
}

//pass an anonymous function as an argument into other function
func PassFunction(i func(p, q string) string) {
	fmt.Println(i("mahbub ", "ahmed"))
}

//unKnown int comes
func proAdder(valus ...int) (int, string) {
	total := 0
	for _, val := range valus {
		total += val
	}
	return total, "nice"
}

func add(v1 int, v2 int) int {
	return v1 + v2
}
func print() {
	fmt.Println("hello")
}
