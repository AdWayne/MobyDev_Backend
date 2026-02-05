package main

import "fmt"

func main() {
	// 1 zadanie
	fmt.Println()
	schooling := "Обучение в школе: 11 years"
	fmt.Println(schooling)
	newschooling := "Новое обучение в школе: 12 years"
	fmt.Println(newschooling)
	fmt.Println()

	// 2 zadanie
	name := "Vladislav"
	fmt.Println(name)
	name = "Azamat"
	fmt.Println(name)
	fmt.Println()

	// 3 zadanie
	var steps int = 0
	fmt.Println("Steps:", steps)
	steps = 2000
	fmt.Println("Обновленный Steps:", steps, "Хорошая работа! Вы уже на пути к своей ежедневной цели.")
	fmt.Println()

	// 4 zadanie
	var largeNumber int = 100000000
	fmt.Println("Large Number:", largeNumber)
	fmt.Println()

	// 5 zadanie
	const breaktime = 15
	fmt.Println("Break time:", breaktime)
	// breaktime = 20 // Это вызовет ошибку компиляции, потому что константы нельзя изменять.
	fmt.Println("Обновденный break time:", breaktime)
}