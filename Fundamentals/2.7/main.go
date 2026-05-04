package main

import ("fmt")

func main()  {
	//1 zadanie
	fmt.Println(5 == 5) //true
	fmt.Println(10 != 3) //true
	fmt.Println(7 > 12) //false
	fmt.Println(15 < 20) //true
	fmt.Println(8 >= 8) //true
	fmt.Println(6 <= 4) //false
	fmt.Println((10 > 5) && (3 < 1)) //false
	fmt.Println((10 > 5) || (3 < 1)) //true
	fmt.Println(!(5 == 5)) //false
	fmt.Println(!(7 < 3)) //true
	fmt.Println(true && false) //false
	fmt.Println(false || false) //false
	fmt.Println(true || false) //true
	fmt.Println((4 + 6 == 10) && (9 > 2)) //true
	fmt.Println((12 / 3 == 4) || (8 < 5)) //true
	fmt.Println()

	//2 zadanie
	var age int = 18
	var hasTicket bool = true
	canEnter := (age >= 18 && hasTicket)
	fmt.Println(canEnter)
	fmt.Println()

	//3 zadanie
	var isLoggedln bool = false
	var isAdmin bool = true
	hasAcces := (isLoggedln && isAdmin) || (isLoggedln && !isAdmin)
	fmt.Println(hasAcces)	
	fmt.Println()
} 