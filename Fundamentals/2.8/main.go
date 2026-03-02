package main

import (
	"fmt"
)

func main() {
	fmt.Println()

	//1 zadanie
	temperature := 8
	if temperature < 0 {
		fmt.Println("Холодно")
	} else if temperature >= 0 && temperature <= 20 {
		fmt.Println("Тепло")
	} else {
		fmt.Println("Жарко")
	}
	fmt.Println()

	//2 zadanie
	score := 45
	if score >= 90 {
		fmt.Println("Отлично")
	} else if score >= 70 && score <= 89 {
		fmt.Println("Хорошо")
	} else if score <= 69 && score >= 50 {
		fmt.Println("Удовлетворительно")
	} else {
		fmt.Println("Не сдал")
	}
	fmt.Println()

	//3 zadanie
	hour := 23
	switch {
	case hour >= 0 && hour <= 5:
		fmt.Println("Ночь")
	case hour >= 6 && hour <= 11:
		fmt.Println("Утро")
	case hour >= 12 && hour <= 17:
		fmt.Println("День")
	case hour >= 18 && hour <= 23:
		fmt.Println("Вечер")
	default:
		fmt.Println("Такое время нету")
	}
	fmt.Println()

	//4 zadanie
	var number int
	fmt.Print("Введите число:")
	fmt.Scanln(&number)
	if number%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
	fmt.Println()

	//5 zadanie
	var day string
	fmt.Print("День:")
	fmt.Scanln(&day)
	fmt.Println()

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Будний день")
	case "Saturday", "Sunday":
		fmt.Println("Выходной")
	default:
		fmt.Println("Некорректный день")
	}
	fmt.Println()

	//6 zadanie
	balance := 5000
	if balance >= 0 {
		fmt.Println("Баланс положительный")
	} else {
		fmt.Println("Баланс отрицательный")
	}
	fmt.Println()

	//7 zadanie
	var age int
	fmt.Print("Возраст:")
	fmt.Scanln(&age)
	if age < 13 {
		fmt.Println("Ребёнок")
	} else if age >= 13 && age <= 17 {
		fmt.Println("Подросток")
	} else {
		fmt.Println("Взрослый")
	}
	fmt.Println()

	//8 zadanie
	var command string
	fmt.Println("У нас есть команды(start, stop, restart):")
	fmt.Scanln(&command)
	switch command {
	case "start":
	fmt.Println("Добро пожоловать и т.д")
	case "restart":
	fmt.Println("Перезагрузка")
	case "stop":
	fmt.Println("До встречи")
	default:
	fmt.Println("Неизвестная команда")
	}
	
	fmt.Println()
	//9 zadanie
	var grade int = 4

	switch grade {
	case 5:
		fmt.Println("Ваша оценка:A")
	case 4:
		fmt.Println("Ваша оценка:B")
	case 3:
		fmt.Println("Ваша оценка:C")
	case 2:
		fmt.Println("Ваша оценка:D")
	case 1:
		fmt.Println("Ваша оценка:F")
	default:
		fmt.Println("У нас 5-балльная система")
	}
}