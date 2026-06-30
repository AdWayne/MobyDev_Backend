package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// CONSTANTS

const (
	BaseDeliveryRate   = 1.0
	MaxBatteryLevel    = 100.0
	EmergencyThreshold = 20.0
	BufferCapacity     = 5
)

// TYPES

type RobptType string

const (
	Drone       RobptType = "Drone"
	Wheeled     RobptType = "Wheeled"
	HeavyLifter RobptType = "HeavyLifter"
)

// STRUCTS

type Product struct {
	ID     int
	Name   string
	Weight float64
}

type Robot struct {
	ID          int
	Model       string
	Type        RobptType
	Battery     float64
	IsAvailable bool
}

// STORAGE

var inventory = map[string]int{}
var robots []Robot

var orderBuffer [BufferCapacity]Product
var orderCount int
var productIDCounter int

var mu sync.Mutex

// INTERFACE

type Mover interface {
	MoveTo(distance float64) error
}

// ROBOT LOGIC

func (r *Robot) MoveTo(distance float64) error {
	var coef float64

	switch r.Type {
	case Drone:
		coef = 1.5
	case Wheeled:
		coef = 0.8
	case HeavyLifter:
		coef = 2.5
	default:
		return fmt.Errorf("неизвестный тип робота")
	}

	energyCost := distance * coef * BaseDeliveryRate

	if r.Battery < EmergencyThreshold {
		return fmt.Errorf("низкий заряд батареи")
	}
	if r.Battery < energyCost {
		return fmt.Errorf("недостаточно энергии")
	}

	r.Battery -= energyCost
	return nil
}

// RECHARGE

func Recharge(robot *Robot) {
	robot.Battery = MaxBatteryLevel
}

// DELIVERY TASK

func DeliveryTask(robot *Robot, product Product, distance float64, reportChan chan<- Report) {
	time.Sleep(time.Duration(distance*100) * time.Millisecond)

	err := robot.MoveTo(distance)

	status := "УСПЕШНО"
	if err != nil {
		status = "НЕУДАЧНО: " + err.Error()
	}

	reportChan <- Report{
		RobotID: robot.ID,
		Text: fmt.Sprintf(
			"Робот %s доставил %s - %s",
			robot.Model,
			product.Name,
			strings.ToUpper(status),
		),
	}
}

// SAFE REPORT STRUCT

type Report struct {
	RobotID int
	Text    string
}

// FIND ROBOT

func findRobotForProduct(product Product) *Robot {
	for i := range robots {
		r := &robots[i]

		if !r.IsAvailable {
			continue
		}

		if r.Battery <= EmergencyThreshold {
			continue
		}

		switch r.Type {
		case Drone:
			if product.Weight > 2.0 {
				continue
			}
		case Wheeled:
			if product.Weight > 20.0 {
				continue
			}
		case HeavyLifter:
			if product.Weight > 100.0 {
				continue
			}
		}

		return r
	}
	return nil
}

// ADD ORDER

func addProductToOrder() {
	if orderCount >= BufferCapacity {
		fmt.Println("Буфер заказов полон")
		return
	}

	var name string
	fmt.Print("Введите название продукта: ")
	fmt.Scan(&name)

	if inventory[name] <= 0 {
		fmt.Println("Товара нет")
		return
	}

	inventory[name]--
	productIDCounter++

	p := Product{
		ID:     productIDCounter,
		Name:   name,
		Weight: 1.0 + float64(productIDCounter%5),
	}

	orderBuffer[orderCount] = p
	orderCount++

	fmt.Println("Добавлено:", p.Name)
}

// SHOW STATE

func showState() {
	fmt.Println("\n=== СКЛАД ===")
	for k, v := range inventory {
		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Println("\n=== РОБОТЫ ===")
	for _, r := range robots {
		status := "свободен"
		if !r.IsAvailable {
			status = "занят"
		}
		fmt.Printf("ID:%d | %s | Battery: %.1f | %s\n",
			r.ID, r.Type, r.Battery, status)
	}

	fmt.Println("\n=== ЗАКАЗЫ ===")
	for i := 0; i < orderCount; i++ {
		p := orderBuffer[i]
		fmt.Printf("ID:%d | %s | %.1f\n", p.ID, p.Name, p.Weight)
	}
}

// DELIVERY ENGINE (FIXED)

func startDelivery() {
	if orderCount == 0 {
		fmt.Println("Нет заказов")
		return
	}

	reportChan := make(chan Report)
	activeJobs := 0

	for i := 0; i < orderCount; i++ {
		product := orderBuffer[i]

		robot := findRobotForProduct(product)
		if robot == nil {
			fmt.Printf("Нет робота для %s\n", product.Name)
			continue
		}

		mu.Lock()
		robot.IsAvailable = false
		mu.Unlock()

		activeJobs++

		go DeliveryTask(robot, product, 2.0+float64(i), reportChan)
	}

	for i := 0; i < activeJobs; i++ {
		report := <-reportChan
		fmt.Println(report.Text)

		mu.Lock()
		for j := range robots {
			if robots[j].ID == report.RobotID {
				robots[j].IsAvailable = true

				if robots[j].Battery < EmergencyThreshold {
					Recharge(&robots[j])
				}
			}
		}
		mu.Unlock()
	}

	orderCount = 0
	close(reportChan)
}

// MAIN

func main() {
	inventory["Коробка сладостей"] = 10
	inventory["Часы"] = 5
	inventory["Телевизор"] = 2

	robots = []Robot{
		{ID: 1, Model: "X1", Type: Drone, Battery: 100, IsAvailable: true},
		{ID: 2, Model: "W1", Type: Wheeled, Battery: 100, IsAvailable: true},
		{ID: 3, Model: "H1", Type: HeavyLifter, Battery: 100, IsAvailable: true},
	}

	for {
		fmt.Println("\n=== МЕНЮ ===")
		fmt.Println("1. Добавить заказ")
		fmt.Println("2. Склад")
		fmt.Println("3. Доставка")
		fmt.Println("4. Выход")

		var choice int
		fmt.Print("Выбор: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addProductToOrder()
		case 2:
			showState()
		case 3:
			startDelivery()
		case 4:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}