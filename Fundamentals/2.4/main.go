package main

import (
	"fmt"
	"math"
)

func main() {
	//1 zadanie
	fmt.Println()
	bannerWidth := 12
	bannerHeight := 8
	bannerArea := bannerHeight * bannerWidth
	halfBannerArea := bannerArea / 2
	bannerBorderLength := (bannerWidth + bannerHeight) * 2
	fmt.Println("Арена площади:", bannerArea)
	fmt.Println("Половина Арена площади:", halfBannerArea)
	fmt.Println("Перемитр Арены:", bannerBorderLength)
	fmt.Println()

	//2 zadanie
	boxCount := 29
	leftoverBoxes := boxCount % 5
	fmt.Println("Остаток от делений:", leftoverBoxes)
	fmt.Println()

	//3 zadanie
	tempMorning := 15
	tempAfternoon := 20
	tempEvening := 30
	totalTemp := tempAfternoon + tempEvening + tempMorning
	averageTemp := totalTemp / 3
	fmt.Println("Тотал:", totalTemp)
	fmt.Println("Авераж:", averageTemp)
	fmt.Println()

	//4 zadanie
	knownWords := 47
	wordsGoal := 120
	progressPercent := (wordsGoal * knownWords) / 100
	fmt.Printf("%d%%", progressPercent)
	fmt.Println()
	fmt.Println()

	//5 zadanie
	coins := 0
	fmt.Println("Начало:", coins)

	coins += 500
	fmt.Println("После задания:", coins)

	coins += 1200
	fmt.Println("После бонуса:", coins)

	coins /= 2
	fmt.Println("После траты половины:", coins)

	coins *= 2
	fmt.Println("После события x2:", coins)

	coins -= 300
	fmt.Println("После финальной траты:", coins)

	//6 zadanie
	var participants int = 42
	var groupCount int = 8
	fmt.Println()

	participantsPerGroup := participants / groupCount
	ostatok := participants % groupCount
	fmt.Println("Целочисленное деление", participantsPerGroup)
	fmt.Println("Остаток:", ostatok)
	fmt.Println()

	//7 zadanie
	fmt.Println(20 - 4*3)     //а тут нету
	fmt.Println((20 - 4) * 3) //ну тут крч скобки перед умножемение
	fmt.Println()

	//8 zadanie
	squareValue := 81
	fmt.Println("Корень 81:", math.Sqrt(float64(squareValue)))

	multiplier := 5
	exponent := 2
	result := math.Pow(float64(multiplier), float64(exponent))
	fmt.Println("Квадрат 5:", result)
}
