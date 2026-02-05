package main
import "fmt"
func main() {
	// 1 zadanie
	fmt.Println()
	var age int = 25
	fmt.Println("Возраст:", age)
	age += 1 // или age = 26
	fmt.Println("Обновленный возраст:", age)
	fmt.Println()

	// 2 zadanie
    var height int = 179
    fmt.Println("Рост:", height)
	height_in_meters := float64(height) / 100.0
	fmt.Println("Рост в метрах:", height_in_meters)
	fmt.Println()

	// 3 zadanie
	var isStudent bool = true
	fmt.Println("Является ли я студентом?:", isStudent)
	fmt.Println()

	// 4 zadanie
	var temperature float64 = 25
	fmt.Println("Температура:", temperature)
	if temperature > 20 {  //По приколу
		fmt.Println("Погода жаркая.")
	} else if temperature < 10 {
		fmt.Println("Погода холодная.")
	}
	fmt.Println()

	// 5 zadanie
	var favoriteQuote string = "Жизнь - это путешествие, а не пункт назначения."
	fmt.Println("Любимая цитата:", favoriteQuote)
	fmt.Println()

	// 6 zadanie
	var PI float64 = 3.14
	fmt.Println("Значение пи:", PI)
	// PI = "3.14159" // Ошибка: нельзя присвоить строку "" переменной типа float64
	// fmt.Println("Обновленное значение пи:", PI)
}