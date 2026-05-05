package main

import "fmt"

func main() {
	const BaseTariff = 0.45
	const HighLoadTax = 0.15
	const NightDiscount = 0.30

	for {
		var name string
		fmt.Print("Введите название прибора (или 'done'): ")
		fmt.Scan(&name)

		if name == "done" {
			fmt.Println("Расчет завершен.")
			break
		}

		var power float64
		var hours float64
		var nightMode bool

		fmt.Print("Мощность (Вт): ")
		fmt.Scan(&power)

		fmt.Print("Время работы (ч): ")
		fmt.Scan(&hours)

		fmt.Print("Ночной режим (true/false): ")
		fmt.Scan(&nightMode)

		consumption := (power * hours) / 1000.0
		cost := consumption * BaseTariff

		if nightMode {
			cost = cost * (1 - NightDiscount)
		}

		if consumption > 10 {
			cost = cost * (1 + HighLoadTax)
		}

		var category string
		switch {
		case power < 100:
			category = "Экономный"
		case power >= 100 && power <= 1000:
			category = "Стандартный"
		default:
			category = "Мощный"
		}

		fmt.Println("\nОтчет по прибору")
		fmt.Printf("Прибор: %s [Категория: %s]\n", name, category)
		fmt.Printf("Расход: %.2f кВт·ч\n", consumption)
		fmt.Printf("Итоговая стоимость: %.2f\n", cost)
	}
}
