package main

import (
	"fmt"
)

func main() {
	//1 zadanie
	for i := 1; i <= 20; i++{
		fmt.Println(i)
	}
	fmt.Println()

	//2 zadanie
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Сумма:",sum)
	fmt.Println()

	//3 zadanie
	number := 5
	for i := 1; i<=10; i++ {
		summa := number * i
		fmt.Println(number,"*", i,"=",summa)
	}
	fmt.Println()

	//4 zadanie
	var n int
	fmt.Print("Введите число:")
	fmt.Scan(&n)
	for i := 1; i<=n; i++ {
		if i%3 == 0{
			fmt.Println(i)
		}
	}
	fmt.Println()

	//5 zadanie
	var number5 int
	fmt.Print("Введите число:")
	fmt.Scan(&number5)
	count := 0

	if number5 == 0 {
		count = 1
	} else {
		for number5 > 0 {
			number5 = number5 / 10
			count++
		}
	}
	fmt.Println("Количество цифр:", count)
	fmt.Println()

	//6 zadanie
	text := "Developer"
	fmt.Println(text)
	for _, i := range text{
		fmt.Println(string(i))
	}

	//7 zadanie
	fmt.Println()
	balance := 3000
	for {
		var vibor int
		fmt.Print("Спрашиваю число:")
		fmt.Scan(&vibor)
		if vibor == 1{
			fmt.Println("Текущий баланс:", balance)
		}else if vibor == 2{
			balance += 500
			fmt.Println("Баланс увеличен. Новый баланс:", balance)
		}else if vibor ==3{
			balance -= 200
			fmt.Println("Баланс уменьшен. Новый баланс:", balance)
		}else if vibor == 0{
			fmt.Println("Пока!")
			break
		}else{
			fmt.Println("Неверный ввод")
		}
	}
} 