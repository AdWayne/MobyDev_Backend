package main

import (
	"fmt"
	"math" // Теперь библиотека используется, и ошибки не будет
)

func main() {
	// var a bool = true
	// var b int = 5
	// var c float64 = 3.1445612354684513546645
	// var d string = "Hi!"
	// fmt.Println("Hello World")
	// fmt.Println("Boolean: ", a)
	// fmt.Println("Integer: ", b)
	// fmt.Printf("Float: %.2f\n", c)
	// fmt.Println("String: ", d)
	// fmt.Print("Rune: ", 'Z')
	// fmt.Println("Я сегодня изучил Go")
	// fmt.Println("Основы Go: это использовать package main and import fmt")
	// fmt.Println("Я человек")

	// age := 18
	// fmt.Println("Мне", age, "лет")
	// age += 1
	// fmt.Println("Через год мне будет", age, "лет")
	// fmt.Println("Сейчас мне", age, "лет")
	
	bannerWidth := 12
	bannerHeight := 8
	bannerArea := bannerWidth * bannerHeight
	fmt.Println("Площадь баннера:", bannerArea)

	halfBannerArea := float64(bannerArea) / 2.0
	fmt.Println("Половина площади баннера:", halfBannerArea)

	bannerBorderLength := 2 * (bannerWidth + bannerHeight)
	fmt.Println("Длина границы баннера:", bannerBorderLength)

	diagonal := math.Sqrt(math.Pow(float64(bannerWidth), 2) + math.Pow(float64(bannerHeight), 2))
	fmt.Println("Диагональ баннера:", diagonal)
}
