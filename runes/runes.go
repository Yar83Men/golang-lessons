package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	start = 48
	end   = 122
)

func main() {
	fmt.Println(generatePassword(15))
}

func generatePassword(count int) string {
	var password strings.Builder
	for i := 1; i <= count; i++ {
		randSymbol := rand.Intn(end-start) + start
		password.WriteString(string(randSymbol))
	}
	return password.String()
}
