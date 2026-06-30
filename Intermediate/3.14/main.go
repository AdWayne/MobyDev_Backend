package main

import (
	"fmt"
)

func sendGreeting(s chan string){
	s <- "Привет из гоутины!"
}

func squareWorker(n int, s chan int) {
	s <- n * n
}

func emitNumbers(n int, s chan int) {
	for i := 1; i <= n; i++ {
		s <- i
	}
	close(s)
}

func sumReader(s chan int, done chan bool) {
	sum := 0

	for num := range s {
		sum += num
	}
	fmt.Println("Сумма:", sum)
	done <- true
}

func filterEven(input chan int, output chan int) {
	for num := range input {
		if num%2 == 0 {
			output <- num
		}
	}
	close(output)
}

func checkChannel(s chan string) {
	val, ok := <-s
	if ok {
		fmt.Println("Получено:", val)
	} else {
		fmt.Println("Канал закрыт")
	}
}

func main() {
	s := make(chan string)
	go sendGreeting(s)
	greeting := <-s
	fmt.Println(greeting)
	fmt.Println()

	// 2
	s2 := make(chan int)
	go squareWorker(9, s2)
	result := <-s2
	fmt.Println("Квадрат числа 9:", result)
	fmt.Println()

	// 3
	s3 := make(chan int)
	go emitNumbers(5, s3)
	fmt.Println("Числа от 1 до 5:")
	for num := range s3 {
		fmt.Println(num)
	}
	fmt.Println()

	// 4
	s4 := make(chan int)
	done := make(chan bool)
	go sumReader(s4, done)
	s4 <- 10
	s4 <- 20
	s4 <- 30
	close(s4)
	<-done
	fmt.Println()

	// 5
	input := make(chan int)
	output := make(chan int)
	go filterEven(input, output)

	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	for num := range output {
		fmt.Println("Чётное число:", num)
	}
	fmt.Println()

	// 6
	s6 := make(chan string)
	go func() {
		s6 <- "Привет!"
	}()
	checkChannel(s6)

	s7 := make(chan string)
	close(s7)
	checkChannel(s7)
	fmt.Println()

	// 7
	s8 := make(chan string, 3)
	s8 <- "Привет!"
	s8 <- "Как дела?"
	s8 <- "Пока!"
	close(s8)
	for msg := range s8 {
	fmt.Println(msg)
}
}