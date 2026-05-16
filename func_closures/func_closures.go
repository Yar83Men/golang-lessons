package main

import "fmt"

func main() {
	counter1 := counter()
	counter2 := counter()

	counter1()
	counter2()
	counter1()
}

func counter() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}
