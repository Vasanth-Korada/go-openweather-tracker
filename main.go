package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type (
	Celsius    float64
	Kelvin     float64
	Fahrenheit float64
)

const (
	WaterBoilingPoint  Celsius = 100
	WaterFreezingPoint Celsius = 0
	AbsoluteZero       Celsius = -273.15
)

func Celsius2Fahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func Celsius2Kelvin(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func Kelvin2Celsius(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func Fahrenheit2Celsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

type APIConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (APIConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return APIConfigData{}, err
	}
	var config APIConfigData

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return APIConfigData{}, err
	}
	return config, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From go\n"))
}

func queryWeather(city string) (WeatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return WeatherData{}, err
	}
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return WeatherData{}, err
	}
	defer response.Body.Close()

	var weatherData WeatherData
	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		return WeatherData{}, nil
	}
	weatherData.Main.Kelvin = float64(Kelvin2Celsius(Kelvin(weatherData.Main.Kelvin)))
	return weatherData, nil
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := queryWeather(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})
	http.ListenAndServe(":8080", nil)
}
