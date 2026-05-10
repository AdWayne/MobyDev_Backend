package main

import (
	"fmt"
)

func main() {
	//1
	var text string
	text = "Golang"
	const author string = "Azamat"
	fmt.Printf("Автор %s написал: %s\n", author, text)
	fmt.Println("Длина текста:", len(text))
	fmt.Println()

	//2
	var message string = "Привет, мир!"
	if message == "" {
		fmt.Println("Строка пустая")
	} else {
		fmt.Println("Строка не пустая")
	}
	fmt.Println("Длина текста:", len(message))
	fmt.Println()

	//3
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scan(&input)
	if len(input) < 5 {
		fmt.Println("Слишком короткое слово")
	} else if len(input) >= 5 && len(input) <= 10 {
		fmt.Println("Нормальная длина")
	} else {
		fmt.Println("Слишком длинное слово")
	}
	fmt.Println()

	//4
	word := "Programma"

	first := word[0]
	last := word[len(word)-1]

	fmt.Printf("Первый символ: %c\n", first)
	fmt.Printf("Последний символ: %c\n", last)

	for i := 0; i < len(word); i++ {
		fmt.Printf("%c\n", word[i])
	}

	fmt.Println("Длина строки:", len(word))

	//5
	sentence := "Hi, salam mister"
	countA := 0

	for _, ch := range sentence {
		if ch == 'a' || ch == 'A' || ch == 'а' || ch == 'А' {
			countA++
		}
	}

	fmt.Printf("В строке %d символов и %d букв a\n",
		len(sentence),
		countA,
	)
	fmt.Println()

	//6
	var soz string

	fmt.Printf("Введите строку: ")
	fmt.Scanln(&soz)

	var reversed string

	for i := len(soz) - 1; i >= 0; i-- {
		reversed += string(soz[i])
	}

	fmt.Println("Исходная строка:", soz)
	fmt.Println("Перевернутая строка:", reversed)

	if soz == reversed {
		fmt.Println("Это палиндром")
	} else {
		fmt.Println("Это не палиндром")
	}
	fmt.Println()

	//8
	var soilem string

	fmt.Print("Введите строку: ")
	fmt.Scanln(&soilem)

	length := len(soilem)

	words := 1
	for i := 0; i < len(soilem); i++ {
		if soilem[i] == ' ' {
			words++
		}
	}

	vowels := 0
	vowelLetters := "aeiouyAEIOUY"

	for i := 0; i < len(soilem); i++ {
		for j := 0; j < len(vowelLetters); j++ {
			if soilem[i] == vowelLetters[j] {
				vowels++
			}
		}
	}

	fmt.Printf(
		"Длина строки: %d, слов: %d, гласных: %d\n",
		length,
		words,
		vowels,
	)
}
