package main

import (
	"fmt"
	"sync"
	"time"
)

func printInfo() {
	fmt.Println("Гоутина запущена")
}

func sayHello(name string) {
	fmt.Println("Привет", name)
}

func heavyTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Выполняю задачу...")
}

func count(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Printf("ID %d -> %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func calculateSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%d * %d = %d\n", num, num, num*num)
}

func checkStatus(site string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Сайт %s доступен\n", site)
}

func processData(id int, category string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Запуск процесса %d в категории %s\n", id, category)
	time.Sleep(delay)
	fmt.Printf("Процесс %d завершен\n", id)
}

func main() {
	//1
	go printInfo()
	time.Sleep(time.Second)
	fmt.Println()

	//2
	names := []string{"Alikhan", "Aisha", "Azamat"}

	for _, name := range names {
		go sayHello(name)
	}
	time.Sleep(time.Second)
	fmt.Println()

	//3
	var wg3 sync.WaitGroup
	wg3.Add(1)
	go heavyTask(&wg3)
	wg3.Wait()
	fmt.Println()

	//4
	var wg4 sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg4.Add(1)
		go count(i, &wg4)
	}

	wg4.Wait()
	fmt.Println()

	//5
	numbers := []int{1, 5, 9, 8}
	var wg5 sync.WaitGroup

	for _, num := range numbers {
		wg5.Add(1)
		go calculateSquare(num, &wg5)
	}

	wg5.Wait()
	fmt.Println()

	//6
	sites := []string{"google.com", "yandex.ru", "go.dev"}
	var wg6 sync.WaitGroup

	for _, site := range sites {
		wg6.Add(1)
		go checkStatus(site, &wg6)
	}

	wg6.Wait()
	fmt.Println()

	//7
	var wg7 sync.WaitGroup

	var category string
	fmt.Print("Введите категорию: ")
	fmt.Scanln(&category)

	delays := []time.Duration{
		500 * time.Microsecond,
		1 * time.Second,
		200 * time.Microsecond,
	}

	for i := 1; i <= 3; i++ {
		wg7.Add(1)
		go processData(i, category, delays[i-1], &wg7)
	}

	wg7.Wait()
	fmt.Println("Все процессы завершены")
}
