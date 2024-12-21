package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка теста, в ней мы определяем expected результат, данные для функции тестирования
	city := "Moskva"
	expected := geo.GeoData{
		City: "Moskva",
	}
	// Act - выполняем функцию, с заданными в моменте Arange данными
	got, err := geo.GetMyLocation(city)
	// Assert - проверка результата работы с expected
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получили %v", expected.City, got.City)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Moscow"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrorNoCity {
		t.Errorf("Ожидалось %v, получение %v", geo.ErrorNoCity, err)
	}
}
