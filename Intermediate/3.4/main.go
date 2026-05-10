package main

import (
	"fmt"
	"sort"
)

func main() {
	//1 zadanie
	var numbers []int
	var x int

	for {
		fmt.Print("0 for exit: ")
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		numbers = append(numbers, x)
	}

	sum := 0

	for _, v := range numbers {
		sum += v
	}
	fmt.Println()
	fmt.Println("Задание 1:")
	fmt.Println("Слайс:", numbers)
	fmt.Println("Сумма:", sum)
	fmt.Println()

	//2 zadanie
	var even []int
	for _, e := range numbers {
		if e%2 == 0 {
			even = append(even, e)
		}
	}
	fmt.Println("Задание 2:")
	fmt.Println("Исходный:", numbers)
	fmt.Println("Четные:", even)
	fmt.Println()

	//3 zadanie
	data := make([]int, len(numbers))
	copy(data, numbers)

	if len(data) > 2 {
		data = append(data[:2], data[3:]...)
	}
	fmt.Println("Задание 3:")
	fmt.Println(data)
	fmt.Println()

	//4 zadanie
	if len(numbers) > 0 {
		min := numbers[0]
		max := numbers[0]

		for _, temps := range numbers {
			if temps < min {
				min = temps
			}
			if temps > max {
				max = temps
			}
		}
		fmt.Println("Задание 4:")
		fmt.Println("Min:", min)
		fmt.Println("Max:", max)
		fmt.Println()
	}

	//5 zadanie
	var words []string
	var str string
	var reversed []string

	for {
		fmt.Print("stop for exit: ")
		fmt.Scan(&str)

		if str == "stop" {
			break
		}

		words = append(words, str)
	}

	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, words[i])
	}

	fmt.Println("Задание 5:")
	fmt.Println("Исходный:", words)
	fmt.Println("Reverse:", reversed)
	fmt.Println()

	//6 zadanie
	sorted := true
	sorted = sort.IntsAreSorted(numbers)
	fmt.Println("Задания 6:", sorted)

	//7 zadanie
	myWishList := []string{"Phone", "Laptop"}
	friendsWishList := []string{"House", "BWV"}

	var registrationList []string

	for _, s := range myWishList{
		registrationList = append(registrationList, s)
	}
	for _, s := range friendsWishList{
		registrationList = append(registrationList, s)
	}

	fmt.Println("Задние 7:",registrationList)

	//8 zadanie
	toyList := []string{"Car", "Doll", "Ball"}
	testToyList := make([]string, len(toyList))
	copy(testToyList, toyList)

	testToyList[1] = "Boat"

	fmt.Println("Задание 8:")
	fmt.Println("toyList:", toyList)
	fmt.Println("testToyList:", testToyList)
}
