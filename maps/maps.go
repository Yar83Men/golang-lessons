package main

import (
	"fmt"
)

// алиасы для типов
type strMap = map[int]string

func main() {
	var map1 = make(strMap, 5)
	map1[1] = "one"
	map1[2] = "two"
	map1[3] = "three"
	map1[4] = "four"
	map1[5] = "five"
	fmt.Printf("%v, len=%d\n", map1, len(map1))

	for key, value := range map1 {
		fmt.Printf("key=%d, value=%s\n", key, value)
	}

	delete(map1, 1)
	fmt.Printf("%v, len=%d\n", map1, len(map1))

	map2 := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println("map2=", map2)
}
