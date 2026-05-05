package main

import "fmt"

//1
func square(n int) int {
	return n * n
}

//2
func maxNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//3
func isEven(n int) bool {
	return n%2 == 0
}

//4
func greetUser(name string) {
	fmt.Println("Привет,", name)
}

//5
func sumSlice(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

//6
func checkLogin(login, password string) bool {
	users := map[string]string{
		"admin": "1234",
	}
	if pass, ok := users[login]; ok && pass == password {
		return true
	}
	return false
}

//7
func increaseBalance(balance *int, amount int) {
	*balance += amount
}

//8
func resetAttempts(attempts *int) {
	if *attempts > 3 {
		*attempts = 0
	}
}

func main() {
	fmt.Println(square(2))
	fmt.Println(square(5))

	fmt.Println(maxNumber(10, 20))
	fmt.Println(maxNumber(7, 3))

	fmt.Println(isEven(4))
	fmt.Println(isEven(7))

	greetUser("Иван")
	greetUser("Алия")

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice(nums))

	fmt.Println(checkLogin("admin", "1234"))
	fmt.Println(checkLogin("user", "0000"))

	balance := 100
	increaseBalance(&balance, 50)
	fmt.Println(balance)

	attempts := 5
	resetAttempts(&attempts)
	fmt.Println(attempts)
}
