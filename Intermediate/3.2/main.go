package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1
	str := "Golang"
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))

	// 2
	str2 := " backend developer "
	fmt.Println(strings.TrimSpace(str2))

	// 3
	str3 := "I like Java"
	contains := strings.Contains(str3, "Java")
	if contains {
		fmt.Println("Найдено")
	} else {
		fmt.Println("Не найдено")
	}

	// 4
	str4 := "one,two,three"
	replace := strings.ReplaceAll(str4, ",", ";")
	fmt.Println(replace)

	// 5
	var str5 string
	fmt.Print("Введите строку: ")
	fmt.Scanf("%s", &str5)
	if str5 == "admin" {
		fmt.Println("Доступ разрешен")
	} else {
		fmt.Println("Доступ запрещен")
	}

	// 6
	str6 := "Go is fast and Go is simple"
	countGo := strings.Count(str6, "Go")
	fmt.Println("Количество:", countGo)

	// 7
	var str7 string
	fmt.Print("Введите строку: ")
	fmt.Scan(&str7)

	str7 = strings.ToLower(str7)
	isPalindrome := true

	for i := 0; i < len(str7)/2; i++ {
		if str7[i] != str7[len(str7)-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("Строка является палиндромом")
	} else {
		fmt.Println("Строка не является палиндромом")
	}

	// 8
	var str8 string
	fmt.Print("Введите строку: ")
	fmt.Scan(&str8)

	found := false

	for i := 0; i < len(str8)-2; i++ {
		if str8[i] == str8[i+1] && str8[i] == str8[i+2] {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Есть три одинаковых символа подряд")
	} else {
		fmt.Println("Нет трёх одинаковых символов подряд")
	}

	// 9
	var password string

	fmt.Print("Введите пароль: ")
	fmt.Scan(&password)

	if len(password) < 8 {
		fmt.Println("Слишком короткий пароль")
	} else {
		hasDigit := false
		hasUpper := false

		for i := 0; i < len(password); i++ {
			if password[i] >= '0' && password[i] <= '9' {
				hasDigit = true
			}

			if password[i] >= 'A' && password[i] <= 'Z' {
				hasUpper = true
			}
		}

		if !hasDigit {
			fmt.Println("Пароль должен содержать цифру")
		} else if !hasUpper {
			fmt.Println("Пароль должен содержать заглавную букву")
		} else {
			fmt.Println("Пароль корректный")
		}
	}

	// 10
	var email string

	fmt.Print("Введите email: ")
	fmt.Scan(&email)

	atCount := 0
	atIndex := -1

	for i := 0; i < len(email); i++ {
		if email[i] == '@' {
			atCount++
			atIndex = i
		}
	}

	if atCount != 1 {
		fmt.Println("Email должен содержать @")
	} else {
		hasDot := false

		for i := atIndex + 1; i < len(email); i++ {
			if email[i] == '.' {
				hasDot = true
				break
			}
		}

		if hasDot {
			fmt.Println("Email корректен")
		} else {
			fmt.Println("Email должен содержать точку после @")
		}
	}

	// 11
	fmt.Print("Введите строку: ")
	fmt.Scan(&str)

	if !(str[0] >= 'A' && str[0] <= 'Z') {
		fmt.Println("Не с заглавной буквы")
	} else if str[len(str)-1] != '.' {
		fmt.Println("Не заканчивается точкой")
	} else {
		fmt.Println("Строка оформлена правильно")
	}

	// 12
	fmt.Print("Введите строку: ")
	fmt.Scan(&str)

	countA := 0
	countE := 0

	for i := 0; i < len(str); i++ {
		if str[i] == 'a' || str[i] == 'A' {
			countA++
		}

		if str[i] == 'e' || str[i] == 'E' {
			countE++
		}
	}

	fmt.Println("Количество букв A/a:", countA)
	fmt.Println("Количество букв E/e:", countE)
}
