package main

import (
	"fmt"
	"go_lessons/http-ex/weather"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	var city string
	color.Green("Введите название города, например 'Moscow'")
	_, err := fmt.Scanln(&city)
	if err != nil {
		color.Red(err.Error())
	}
	city = strings.ToUpper(city[:1]) + strings.ToLower(city[1:])

	weatherResponse, err := weather.GetWeather(city)
	if err != nil {
		color.Red("Ошибка определения локации")
		time.Sleep(3000 * time.Millisecond)
		return
	}

	if weatherResponse == nil {
		color.Red("Ошибка определения локации")
		time.Sleep(3000 * time.Millisecond)
		return
	}
	var dayOrNight string
	if weatherResponse.Current.IsDay == 0 {
		dayOrNight = "ночь"
	} else {
		dayOrNight = "день"
	}

	finalResponse := "Погда в городе:" + city + "\n" +
		"Температура:" + strconv.FormatFloat(weatherResponse.Current.TempC, 'g', -1, 64) + "\n" +
		"Время суток:" + dayOrNight + "\n" +
		"Давление:" + strconv.FormatFloat(weatherResponse.Current.PressureMb, 'g', -1, 64) + "\n" +
		"Вероятность облачности:" + strconv.Itoa(weatherResponse.Current.Cloud) + "\n" +
		"Вероятность дождя:" + strconv.Itoa(weatherResponse.Current.WillItRain) + "\n"
	color.Green(finalResponse)
	color.Cyan("По истечению 7-ми сек. закроется")
	time.Sleep(7000 * time.Millisecond)
}
