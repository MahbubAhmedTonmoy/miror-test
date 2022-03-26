package main

import (
	"fmt"
)

type X interface {
	int | float32 | float64 | ~string
}

func sum[T X](v1, v2 T) T {
	return v1 + v2
}

func max[T X](input []T) T {
	var max T
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	return max
}
func min[T X](a T, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	fmt.Println(sum[int](1, 2))
	fmt.Println(sum(1.5, 2.5))
	fmt.Println(sum("1", "2"))
	fmt.Println(max([]int{1, 2, 3, 4}))
	fmt.Println(max([]float32{1.5, 2.5}))
	fmt.Println(max([]string{"1", "2"}))

	fmt.Println(min(1, 2))
	fmt.Println(min(1.5, 2.5))
	fmt.Println(min("f", "c"))
}
