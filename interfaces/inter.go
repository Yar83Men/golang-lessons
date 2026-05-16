package main

import "github.com/fatih/color"

func main() {
	PrintErrors("Hello")
	UnpackInterface(123)
}

func PrintErrors(value any) {
	switch valueType := value.(type) {
	case string:
		color.Cyan(valueType)
	case int:
		color.Red("Код ошибки %d", valueType)
	case error:
		color.Red(valueType.Error())
	default:
		color.Green("Неизвестный код ошибки")
	}
}

func UnpackInterface(value any) {
	val, ok := value.(int)
	if ok {
		color.Green("Число %d", val)
	}
}
