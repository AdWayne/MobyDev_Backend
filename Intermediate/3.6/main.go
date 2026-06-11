package main

import (
	"fmt"
	"reflect"
)

func main() {
	//1 zadanie
	toolUsage := map[string]int{"Go": 3, "VSCode": 5, "Git": 2}
	fmt.Println("Задание 1:")
	for tool, count := range toolUsage {
		fmt.Println(tool, "=", count)
	}
	fmt.Println()

	//2 zadanie
	buildStatus := map[string]bool{"build": true, "run": false}
	fmt.Println("Задание 2:")
	if buildStatus["build"] {
		fmt.Println("Сборка прошла успешна")
	}
	fmt.Println()

	//3 zadanie
	var name string
	fmt.Println("Задание 3:")
	fmt.Println("Введите имя:")
	fmt.Scan(&name)
	userInfo := map[string]interface{}{"name": name, "isLoggedIn": true}
	fmt.Println("Пользователь", userInfo["name"], "вошел в систему!")
	fmt.Println()

	//4 zadanie
	fmt.Println("Задание 4:")
	cpuLoad := map[int]int{1: 40, 2: 65, 3: 30}

	maxCore := 0
	maxLoad := 0

	for core, load := range cpuLoad {
		if load > maxLoad {
			maxLoad = load
			maxCore = core
		}
	}
	fmt.Println("Максимальная загрузка ядра:", maxCore)
	fmt.Println()

	//5 zadanie
	fmt.Println("Задание 5:")
	examResults := map[string]int{"Aruzhan": 85, "Dias": 92, "Alina": 78}
	for name, grades := range examResults {
		if grades >= 80 {
			fmt.Println(name, "сдал экзамен")
		} else {
			fmt.Println(name, "не сдал экзамен")
		}
	}
	fmt.Println()

	//6 zadanie
	fmt.Println("Задание 6:")
	words := []string{"go", "is", "fast"}
	wordLength := make(map[string]int)

	for _, w := range words {
		wordLength[w] = len(w)
	}
	fmt.Println(wordLength)
	fmt.Println()

	//7 zadanie
	fmt.Println("Задание 7:")
	menu := map[string]int{"Burger": 2500, "Pizza": 3200, "Tea": 500}

	var food string
	fmt.Println("Введите блюдо:")
	fmt.Scan(&food)

	if price, ok := menu[food]; ok {
		fmt.Println("Цена:", price)
	} else {
		fmt.Println("Блюда не найдено")
	}
	fmt.Println()

	//8 zadanie!
	fmt.Println("Задание 8:")
	loginAttempts := map[string]int{"admin": 2, "guest": 0}

	loginAttempts["admin"]++

	if loginAttempts["admin"] > 2 {
		fmt.Println("Доступ заблокирован")
	}
	fmt.Println()

	//9 zadanie
	fmt.Println("Задание 9:")
	sales := [2][3]int{
		{10, 20, 30},
		{5, 15, 25},
	}

	total := make(map[int]int)

	for store := 0; store < 2; store++ {
		sum := 0
		for day := 0; day < 3; day++ {
			sum += sales[store][day]
		}
		total[store+1] = sum
	}

	fmt.Println(total)
	fmt.Println()

	//10 zadanie
	fmt.Println("Задание 10:")
	numbers := []int{4, 7, 2, 9, 5}
	numberStatus := make(map[int]string)

	for _, n := range numbers {
		if n%2 == 0 {
			numberStatus[n] = "even"
		} else {
			numberStatus[n] = "odd"
		}
	}
	fmt.Println(numberStatus)
	fmt.Println()

	//11 zadanie
	fmt.Println("Задание 11:")
	defaultConfig := map[string]string{"host": "localhost", "port": "8080", "mode": "production"}
	currentConfig := map[string]string{"host": "localhost", "port": "8080", "mode": "production"}

	if reflect.DeepEqual(defaultConfig, currentConfig) {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}
	//Изменим Конфигурацию
	currentConfig["mode"] = "debug"

	if reflect.DeepEqual(defaultConfig, currentConfig) {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}
}
