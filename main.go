package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Новый проект")
	// в методе flag.String("название поля", "Значение по-умолчанию", "Пояснение"), в результате получаем указатель на строку
	city := flag.String("city", "Moscow", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	// Используется только после того, как мы объявили все необходимые нам аргументы, то есть в конце
	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
