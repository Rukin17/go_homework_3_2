package main

import (
	"fmt"
	"go_homework_3_2/data"
)

func main() {
	fmt.Printf("База городов: %v\n\n", data.Cities)

	city := "Сочи"

	// add
	data.AddCity(city)
	fmt.Printf("Город %s - Успешно добавлен!\n%v\n\n", city, data.Cities)

	// Search
	if data.SearchCity(city) {
		fmt.Printf("Город %s - Найден!\n\n", city)
	} else {
		fmt.Printf("Город %s - Не найден!\n\n", city)
	}

	// Remove
	if data.RemoveCity(city) {
		fmt.Printf("Город %s - Удален из списка!\n %v\n", city, data.Cities)
	} else {
		fmt.Printf("Город %s - Не найден!\n\n", city)
	}
}
