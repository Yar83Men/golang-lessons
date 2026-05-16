package main

import "fmt"

func main() {
	slice1 := make([]int, 0)
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	slice1 = append(slice1, 3)
	slice2 := append(slice1, 4)
	slice3 := append(slice1, 5, 6)

	fmt.Println("Slice1=", slice1)
	fmt.Println("Slice2=", slice2)
	fmt.Println("Slice3=", slice3)
}
