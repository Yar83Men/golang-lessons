package main

import "fmt"

func main() {
	fmt.Println(sum(100, 100.55))
}

func sum[T int | int8 | int16 | int32 | float32 | float64](a, b T) T {
	return a + b
}

type List[T any] struct {
	elements []T
}
