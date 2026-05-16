package main

import "fmt"

func main() {
	fmt.Println("Массивы")
	transactions := []int{1, 100, 200}
	var bancks []int = []int{300, 500}

	fmt.Printf("%v, %T \n", transactions, transactions)
	fmt.Printf("%v, %T \n", bancks, transactions)

	slice := transactions[:]
	fmt.Printf("Slice %v, %T \n", slice, slice)
	slice = append(slice, 100, 500)
	fmt.Printf("Slice %v, %T \n", slice, slice)
	fmt.Println("transactions cap and len", cap(transactions), len(transactions))
	fmt.Println("slice cap and len", cap(slice), len(slice))
	fmt.Println("---------------------------------------------------------------")
	slice1 := []string{}
	slice1 = append(slice1, "hello", "go")
	fmt.Printf("Slice1=%v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, "Java")
	fmt.Printf("Slice1=%v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, "Python", "C#")
	fmt.Printf("Slice1=%v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, "Lua", "JavaScript", "CoffeScript")
	fmt.Printf("Slice1=%v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, "PHP", "Perl", "Bash")
	fmt.Printf("Slice1=%v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
}
