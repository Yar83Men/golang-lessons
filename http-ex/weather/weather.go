package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/fatih/color"
)

const WEATHER_URI = "https://api.weatherapi.com/v1/current.json"
const API_KEY = "919b0eb97dfc4906963192308261505"

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Current struct {
	TempC      float64 `json:"temp_c"`
	IsDay      int     `json:"is_day"`
	PressureMb float64 `json:"pressure_mb"`
	Cloud      int     `json:"cloud"`
	WillItRain int     `json:"will_it_rain"`
}

type Location struct {
	Name string `json:"name"`
}

func GetWeather(city string) (*Weather, error) {
	baseUrl, err := url.Parse(WEATHER_URI)
	if err != nil {
		fmt.Println(err.Error())
	}
	params := url.Values{}
	params.Add("key", API_KEY)
	params.Add("q", city)
	params.Add("aqi", "no")
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("ошибка получения даных о погоде")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("ошибка парсинга ответа")
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		color.Red(err.Error())
		return nil, errors.New("ошибка парсинга ответа")
	}

	if weather.Location.Name == "" {
		return nil, errors.New("локация не найдена")
	}

	return &weather, nil
}
