package crud

import (
	"go_homework_3_2/db"
	"strings"
)

func AddCity(cityName string) {
	db.Cities = append(db.Cities, cityName)
}

func SearchCity(cityName string) bool {
	cityName = strings.ToLower(cityName)
	flag := false
	for _, city := range db.Cities {
		if cityName == strings.ToLower(city) {
			flag = true
		}
	}
	return flag
}

func RemoveCity(cityName string) bool {
	cityName = strings.ToLower(cityName)
	flag := false
	for i, city := range db.Cities {
		if cityName == strings.ToLower(city) {
			db.Cities = append(db.Cities[:i], db.Cities[i+1:]...)
			flag = true
		}
	}
	return flag
}
