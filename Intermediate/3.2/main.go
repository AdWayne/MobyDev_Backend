package main

import (
	"fmt"
	"strings"
)

func main() {
	// Задание 1
	var text string = "Программирование на Go"
	const author string = "Azamat"

	fmt.Printf("Автор %s написал: %s\n", author, text)
	fmt.Println("Длина текста:", len(text))
	fmt.Println()

	// Задание 2
	var message string = "Я учу Golang язык программирование"

	if message == "" {
		fmt.Println("Строка пуста")
	} else {
		fmt.Println("Строка не пустая")
	}

	fmt.Println("Длина строки:", len(message))
	fmt.Println()

	// Задание 3
	var input string
	fmt.Print("Введите слово: ")
	fmt.Scan(&input)

	if len(input) < 5 {
		fmt.Println("Слишком короткое слово")
	} else if len(input) <= 10 {
		fmt.Println("Нормальная длина")
	} else {
		fmt.Println("Слишком длинное слово")
	}
	fmt.Println()

	// Задание 4
	var word string = "Golang"

	fmt.Println("Первый символ:", string(word[0]))
	fmt.Println("Последний символ:", string(word[len(word)-1]))

	fmt.Println("Все символы:")
	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))
	}

	fmt.Println("Длина слова:", len(word))
	fmt.Println()

	// Задание 5
	var sentence string = "Мама"

	countA := 0
	for _, ch := range sentence {
		if ch == 'а' || ch == 'А' {
			countA++
		}
	}

	fmt.Printf("В строке %d символов и %d букв а\n", len(sentence), countA)
	fmt.Println()

	// Задание 6
	var soz string
	fmt.Print("Проверка палиндрома: ")
	fmt.Scan(&soz)

	var reversed string
	for i := len(soz) - 1; i >= 0; i-- {
		reversed += string(soz[i])
	}

	if soz == reversed {
		fmt.Println("Это палиндром")
	} else {
		fmt.Println("Это не палиндром")
	}

	fmt.Println("Исходная строка:", soz)
	fmt.Println("Перевёрнутая строка:", reversed)
	fmt.Println()

	// Задание 7
	var line string
	fmt.Print("Введите текст: ")
	fmt.Scanln()
	fmt.Scanln(&line)

	length := len(line)
	words := len(strings.Fields(line))

	vowels := "aeiouyAEIOUY"
	countVowels := 0

	for i := 0; i < len(line); i++ {
		if strings.ContainsRune(vowels, rune(line[i])) {
			countVowels++
		}
	}

	fmt.Printf("Длина строки: %d, слов: %d, гласных: %d\n", length, words, countVowels)
}
