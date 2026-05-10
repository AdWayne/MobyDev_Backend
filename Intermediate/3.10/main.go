package main

import (
	"fmt"
	"math"
)

type Person struct {
	name string
	age  int
}

type Book struct {
	title  string
	author string
	pages  int
}

type Rectangle struct {
	width  float64
	height float64
}

type Car struct {
	brand string
	year  int
}

type Student struct {
	name  string
	grade int
}

type Circle struct {
	radius float64
}

type Car7 struct {
	brand string
	speed int
}

func CheckSpeed(car Car7) string {
	if car.speed > 100 {
		return "Слишком быстро"
	} else if car.speed >= 60 && car.speed <= 100 {
		return "Нормально"
	} else {
		return "Медленно"
	}
}

func main() {
	//1
	Azamat := Person{"Azamat", 19}
	fmt.Println(Azamat)
	fmt.Println()

	//2
	Kniga := Book{"Дневник памяти", "Николас Спаркс", 240}
	fmt.Println(Kniga)
	fmt.Println()

	//3
	car := Car{
		brand: "Toyota",
		year:  2007,
	}

	carpointer := &car
	carpointer.year = 2022
	fmt.Println("Марка:", car.brand)
	fmt.Println("Год:", car.year)
	fmt.Println()

	//4
	rect := Rectangle{
		width:  5,
		height: 3,
	}

	S := rect.width * rect.height
	P := 2 * (rect.width + rect.height)
	fmt.Println("Площадь", S)
	fmt.Println("Периметр", P)

	rect.width = 10
	rect.height = 8

	newS := rect.width * rect.height
	newP := 2 * (rect.width + rect.height)
	fmt.Println()
	fmt.Println("Новый Площадь", newS)
	fmt.Println("Новый Периметр", newP)
	fmt.Println()

	//5
	student1 := Student{
		name:  "Azamat",
		grade: 99,
	}

	student2 := Student{
		name:  "Aruzhan",
		grade: 85,
	}

	fmt.Printf("Студент %s получил %d баллов\n", student1.name, student1.grade)
	fmt.Printf("Студент %s получил %d баллов\n", student2.name, student2.grade)
	fmt.Println()
	if student1.grade > student2.grade {
		fmt.Println("Больше баллов получил:", student1.name)
	} else {
		fmt.Println("Больше баллов получил:", student2.name)
	}
	fmt.Println()

	//6
	circle := Circle{
		radius: 5,
	}

	circleS := math.Pi * circle.radius * circle.radius
	circleL := 2 * math.Pi * circle.radius

	fmt.Println("Площадь круга:", circleS)
	fmt.Println("Длина окружности:", circleL)
	fmt.Println()
	circle.radius = 10

	circlenewS := math.Pi * circle.radius * circle.radius
	circlenewL := 2 * math.Pi * circle.radius

	fmt.Println("New Площадь круга:", circlenewS)
	fmt.Println("New Длина окружности:", circlenewL)
	fmt.Println()

	//7
	car1 := Car7{
		brand: "BMW",
		speed: 120,
	}
	car2 := Car7{
		brand: "Toyota",
		speed: 60,
	}
	car3 := Car7{
		brand: "Mercedez",
		speed: 90,
	}

	fmt.Printf("Машина %s: %s\n", car1.brand, CheckSpeed(car1))
	fmt.Printf("Машина %s: %s\n", car2.brand, CheckSpeed(car2))
	fmt.Printf("Машина %s: %s\n", car3.brand, CheckSpeed(car3))
}
