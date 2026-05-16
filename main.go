package main

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	statusError = 500
	statusOk    = 200
)

func main() {
	var inputAge int
	var inputName string
	var endInput string
	var enterPush int
	//fmt.Println("Введите возраст")
	color.Green("Введите возраст:")
	var _, errAge = fmt.Scan(&inputAge)
	if errAge != nil {
		color.Red("Ошибка введения возраста")
		//fmt.Println("Ошибка введения возраста")
	}
	color.Green("Введите имя")
	// fmt.Println("Введите имя")
	var _, errName = fmt.Scan(&inputName)
	if errName != nil {
		color.Red("Ошибка введения имени")
		//fmt.Println("Ошибка введения имени")
	}
	//fmt.Printf("Ваше имя=%s, возраст=%d\n", inputName, inputAge)
	color.Yellow("Ваше имя=%s, возраст=%d\n", inputName, inputAge)
	color.Red("Введите любой символ для завершения")
	//fmt.Println("Введите любой символ для завершения")
	enterPush, _ = fmt.Scan(&endInput)
	if enterPush >= 1 {
		return
	}
}
