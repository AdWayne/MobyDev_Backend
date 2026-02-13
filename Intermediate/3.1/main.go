package main

import "fmt"

func main() {
	fmt.Println()
	//1 zadanie
	run := [3]string{"Бег","Разножка","Олений бег"}
	fmt.Println(run)
	hodba := [2]string{"Ходьба","Шаги"}
	fmt.Println(hodba)
	lenght := len(run)
	fmt.Println(lenght)
	lenght2 := len(hodba)
	fmt.Println(lenght2)
	fmt.Println()
	
	//2 zadanie
	subjectsList := [3]string{"Физика","Химия","География"}
	fmt.Println(subjectsList[0])
	fmt.Println(subjectsList[2])
	
	subjectsList[1] = "Биология"

	fmt.Println("Обновленный массив:", subjectsList)
	fmt.Println()

	//3 zadanie
	Person := [3]string{"Tom","35","New York"}
	name := Person[0]
	age := Person[1]
	home := Person[2]

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(home)
	fmt.Println()

	//4 zadanie
	numberList := [5]int{1,2,3,4,5}
	found := false

	for i := 0; i < len(numberList); i++ {
		if numberList[i] == 3{
			found = true
			break
		}
	}
	if found{
		fmt.Println("Нашел")
	}else {
		fmt.Println("Не нашел")
	}
	fmt.Println()

	//5 zadanie
	friendsList := [3]string{"Medet","Bahytjan","Alpamus"}
	found5 := false

	for _, name := range friendsList {
		if name == "Bekbolat" {
			found5 = true
			break
		}
	}

	if found5 {
		fmt.Println("Мне очень повезло")
	}else{
		fmt.Println("Мне не повезло")
	}
	fmt.Println()

	//6 zadanie
	firstList := [3]int{1,2,3}
	secondList := [3]int{1,2,4}
	
	if firstList == secondList{
		fmt.Println("Массивы верны")
	}else {
		fmt.Println("Не равны")
	}
	fmt.Println()

	//7 zadanie
	myWishList := []string{"Квартира в Алмате,"}
	friendsWishList := []string{"Альтушку,","IT Startapp."}
	registrationList := []string{}

	for _, item := range myWishList {
		registrationList = append(registrationList, item)
	}

	for _, item := range friendsWishList{
		registrationList = append(registrationList, item)
	}
	fmt.Println(myWishList)
	fmt.Println(friendsWishList)
	fmt.Println("Общий список подарков:", registrationList)
	fmt.Println()

	//8 zadanie
	toyList := [3]string{"Car","Doll","Ball"}
	testToyList := toyList

	testToyList[1] = "Boat"
	fmt.Println(toyList)
	fmt.Println(testToyList)
}