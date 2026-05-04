package main

import (
	"fmt"
	"strings"
)

func main() {
	menu := map[string]float64{
		"Эспрессо": 800.0,
		"Латте":    1200.0,
		"Капучино": 1100.0,
		"Сэндвич":  1500.0,
		"Круассан": 900.0,
	}

	order := []string{}

	fmt.Println("Добро пожаловать в наше кафе!")
	fmt.Println("Меню:")
	for item, price := range menu {
		fmt.Printf("- %s: %.2f тг\n", item, price)
	}
	fmt.Println("Ваше блюдо для заказа или exit для завершения:")

	for {
		var input string
		fmt.Scan(&input)

		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			break
		}

		if _, exists := menu[input]; exists {
			order = append(order, input)
			fmt.Printf("Добавлено: %s\n", input)
		} else {
			fmt.Println("К сожалению нету такого блюдо")
		}
	}

	sum := make(map[string]int)
	var total float64

	for _, item := range order {
		sum[item]++
		total += menu[item]
	}

	var discount float64
	if total > 5000 {
		discount = total * 0.10
	}

	tax := (total - discount) * 0.12
	finaltotal := (total - discount) + tax

	fmt.Println("\nФинальный чек")
	if len(order) == 0 {
		fmt.Println("Корзина пуста")
	} else {
		for item, count := range sum {
			price := menu[item] * float64(count)
			fmt.Printf("%s x%d — %.2f тг\n", item, count, price)
		}

		fmt.Println("---------------------")
		fmt.Printf("Сумма без скидки: %.2f тг\n", total)
		if discount > 0 {
			fmt.Printf("Скидка (10%%): -%.2f тг\n", discount)
		}
		fmt.Printf("НДС (12%%): %.2f тг\n", tax)
		fmt.Printf("ИТОГО К ОПЛАТЕ: %.2f тг\n", finaltotal)
	}
	fmt.Println("---------------------")
}
