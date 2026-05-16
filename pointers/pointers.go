package pointers

import "fmt"

func main() {
	slice := make([]int, 2, 3)
	pointer := &slice
	fmt.Printf("%v %T\n", *pointer, pointer)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%p \n", &arr)
	arrSlice := arr[:]
	fmt.Printf("%p \n", &arrSlice)
	// revesred := reverseArray(arrSlice)
	ReversedSlice(arrSlice)
	fmt.Println("Slice=", arrSlice)
	fmt.Println("Array=", arr)
	fmt.Printf("%p \n", &arr)
	fmt.Printf("%p \n", &arrSlice)
}

func reverseArray(arr []int) []int {
	reversed := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}

	return reversed
}

func ReversedSlice(arr []int) {
	arrayLen := len(arr)
	for i := range arrayLen / 2 {
		j := arrayLen - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}
