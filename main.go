package main

import (
	"fmt"
	"go_homework_3_2/crud"
	"go_homework_3_2/db"
)

func main() {
	fmt.Printf("База городов: %v\n\n", db.Cities)

	city := "Сочи"

	// add
	crud.AddCity(city)
	fmt.Printf("Город %s - Успешно добавлен!\n%v\n\n", city, db.Cities)

	// Search
	if crud.SearchCity(city) {
		fmt.Printf("Город %s - Найден!\n\n", city)
	} else {
		fmt.Printf("Город %s - Не найден!\n\n", city)
	}

	// Remove
	if crud.RemoveCity(city) {
		fmt.Printf("Город %s - Удален из списка!\n %v\n", city, db.Cities)
	} else {
		fmt.Printf("Город %s - Не найден!\n\n", city)
	}
}
