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
	city, err := getCity()
	weatherResponse, err := weather.GetWeather(city)
	if err != nil {
		errorHandler()
		return
	}
	if weatherResponse == nil {
		errorHandler()
		return
	}
	color.Yellow("Погда в городе:" + city + "\n")
	color.Yellow("Температура:" + getTemperature(weatherResponse.Current.TempC) + "\n")
	color.Magenta("Время суток:" + timesOfDay(weatherResponse.Current.IsDay) + "\n")
	color.Red("Давление:" + strconv.FormatFloat(weatherResponse.Current.PressureMb, 'g', -1, 64) + "\n")
	color.Cyan("Вероятность облачности:" + strconv.Itoa(weatherResponse.Current.Cloud) + "\n")
	color.Blue("Вероятность дождя:" + strconv.Itoa(weatherResponse.Current.WillItRain) + "\n")

	color.Cyan("По истечению 7-ми сек. закроется")
	time.Sleep(7000 * time.Millisecond)
}

func getCity() (string, error) {
	var city string
	color.Green("Введите название города, например 'Moscow' или 'Eupatoria'")
	_, err := fmt.Scanln(&city)
	if err != nil {
		color.Red(err.Error())
	}
	city = strings.ToUpper(city[:1]) + strings.ToLower(city[1:])
	return city, err
}

func errorHandler() {
	color.Red("Ошибка определения локации, зашел не в ту дверь)")
	time.Sleep(3000 * time.Millisecond)
}

func timesOfDay(isDay int) string {
	var dayOrNight string
	if isDay == 0 {
		dayOrNight = "Мойте ноги ложитесь спать)"
	} else {
		dayOrNight = "Солнце светит негры пашут)"
	}
	return dayOrNight
}

func getTemperature(temp float64) string {
	var result string
	temperatureStr := strconv.FormatFloat(temp, 'g', -1, 64)
	if temp < 20 {
		result += "Дубачина, как в холодильнике " + temperatureStr + "\n"
	} else if temp > 20 && temp < 30 {
		result += "Прохладно, норм " + temperatureStr + "\n"
	} else {
		result += "Шкварит как в микроволновке " + temperatureStr + "\n"
	}
	return result
}
