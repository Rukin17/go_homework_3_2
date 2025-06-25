package data

import (
	"strings"
)

func AddCity(cityName string) {
	Cities = append(Cities, cityName)
}

func SearchCity(cityName string) bool {
	cityName = strings.ToLower(cityName)
	flag := false
	for _, city := range Cities {
		if cityName == strings.ToLower(city) {
			flag = true
		}
	}
	return flag
}

func RemoveCity(cityName string) bool {
	cityName = strings.ToLower(cityName)
	flag := false
	for i, city := range Cities {
		if cityName == strings.ToLower(city) {
			Cities = append(Cities[:i], Cities[i+1:]...)
			flag = true
		}
	}
	return flag
}
