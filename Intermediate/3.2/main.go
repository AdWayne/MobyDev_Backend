package main

import "fmt"

func main() {
	//1 zadanie
	var massiv [3][5]int
	row, col := 0, 0
	maxVal := -1 << 63 //для отрицательных м максимальных значений
	fmt.Println("Массив 3 столбец 5 строк")
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&massiv[i][j])
			if massiv[i][j] > maxVal {
				maxVal = massiv[i][j]
				row, col = i, j
			}
		}
	}
	fmt.Println(row, col)
	fmt.Println()

	//2 zadanie
	const a = 11
	var massiv2 [a][a]string
	mid := a / 2

	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			if i == mid || j == mid || i == j || i+j == a-1 {
				massiv2[i][j] = "*"
			} else {
				massiv2[i][j] = "."
			}
		}
	}

	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			fmt.Print(massiv2[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()

	//3 zadanie
	const c = 8
	var chessporidak [c][c]string
	for i := 0; i < c; i++ {
		for j := 0; j < c; j++ {
			if (i+j)%2 == 0 {
				chessporidak[i][j] = "."
			} else {
				chessporidak[i][j] = "*"
			}
			fmt.Print(chessporidak[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()

	//4-5 zadanie
	var matrix [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	var i, j int
	fmt.Scan(&i, &j)

	//4
	matrix[i-1], matrix[j-1] = matrix[j-1], matrix[i-1]

	for _, row := range matrix {
	    fmt.Println(row)
	}

	//5
	// for row := 0; row < 4; row++ {
	// 	matrix[row][i], matrix[row][j] = matrix[row][j], matrix[row][i]
	// }

	// for r := 0; r < 4; r++ {
	// 	for c := 0; c < 4; c++ {
	// 		fmt.Printf("%d ", matrix[r][c])
	// 	}
	// 	fmt.Println()
	// }
}
