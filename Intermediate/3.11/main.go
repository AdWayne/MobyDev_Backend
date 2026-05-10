package main

import (
	"fmt"
	"strings"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Reactangle struct {
	Width  float64
	Height float64
}

func (r Reactangle) Area() float64 {
	return r.Width * r.Height
}

func (r Reactangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Student struct {
	Name   string
	Grades []int
}

type Person interface {
	Average() float64
}

func (s Student) Average() float64 {
	sum := 0

	for _, grade := range s.Grades {
		sum += grade
	}

	return float64(sum) / float64(len(s.Grades))
}

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64)
}

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {
	if b.Balance >= amount {
		b.Balance -= amount
	} else {
		fmt.Println("Недостаточно средств у", b.Owner)
	}
}

type Analyzer interface {
	CountWords() int
	CountSentences() int
	CountVowels() int
}

type TextAnalyzer struct {
	Text string
}

func (t TextAnalyzer) CountWords() int {
	words := strings.Fields(t.Text)
	return len(words)
}

func (t TextAnalyzer) CountSentences() int {
	count := 0

	for _, char := range t.Text {
		if char == '.' || char == '!' || char == '?' {
			count++
		}
	}

	return count
}

func (t TextAnalyzer) CountVowels() int {
	count := 0

	vowels := "aeiouAEIOUаәеёиіоуүұыэюяAEIOUАӘЕЁИІОУҮҰЫЭЮЯ"

	for _, char := range t.Text {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
func main() {
	//1
	rect1 := Reactangle{
		Width:  10,
		Height: 4,
	}

	rect2 := Reactangle{
		Width:  5,
		Height: 3,
	}

	shapes := []Shape{rect1, rect2}

	for i, shape := range shapes {
		fmt.Printf(
			"Прямоугольник %d: площадь %.2f, периметр %.2f\n",
			i+1,
			shape.Area(),
			shape.Perimeter(),
		)
	}
	fmt.Println()

	//2
	student1 := Student{
		Name:   "Azamat",
		Grades: []int{90, 85, 75},
	}
	student2 := Student{
		Name:   "Aruzhan",
		Grades: []int{90, 82, 70},
	}
	student3 := Student{
		Name:   "Aisha",
		Grades: []int{100, 100, 100},
	}

	people := []Person{student1, student2, student3}
	students := []Student{student1, student2, student3}

	for i, person := range people {
		fmt.Printf(
			"Студент %s: средний балл %.2f\n",
			students[i].Name,
			person.Average(),
		)
	}
	fmt.Println()

	//3
	account1 := BankAccount{
		Owner:   "Azamat",
		Balance: 1000,
	}
	account2 := BankAccount{
		Owner:   "Aisha",
		Balance: 10000,
	}

	account1.Deposit(3000)
	account1.Withdraw(500)

	account2.Deposit(5000)
	account2.Withdraw(400)

	accounts := []Account{&account1, &account2}

	for _, account := range accounts {
		if acc, ok := account.(*BankAccount); ok {
			fmt.Printf(
				"Счёт %s: баланс %.2f\n",
				acc.Owner,
				acc.Balance,
			)
		}
	}
	fmt.Println()

	//4
	text := TextAnalyzer{
		Text: "Привет! Че как там дела? Я привесттвую вас.",
	}

	var analyzer Analyzer = text
	fmt.Println(text)
	fmt.Printf(
		"Слов: %d, предложений: %d, гласных: %d\n",
		analyzer.CountWords(),
		analyzer.CountSentences(),
		analyzer.CountVowels(),
	)
}
