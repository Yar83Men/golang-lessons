package main

import (
	"flag"
	"fmt"
)

func main() {
	flag1 := flag.String("flag1", "", "")
	fmt.Printf("flag1: %s\n", *flag1)
}
