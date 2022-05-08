package testing

import (
	"fmt"
	"sudoku/game"
)

// var sample2 [81]int
// var solutn2 [81]int

func sample1() [81]int {
	return [81]int{
		5, 3, 0, 0, 7, 0, 0, 0, 0,
		6, 0, 0, 1, 9, 5, 0, 0, 0,
		0, 9, 8, 0, 0, 0, 0, 6, 0,
		8, 0, 0, 0, 6, 0, 0, 0, 3,
		4, 0, 0, 8, 0, 3, 0, 0, 1,
		7, 0, 0, 0, 2, 0, 0, 0, 6,
		0, 6, 0, 0, 0, 0, 2, 8, 0,
		0, 0, 0, 4, 1, 9, 0, 0, 5,
		0, 0, 0, 0, 8, 0, 0, 7, 9,
	}
}

func solutn1() [81]int {
	return [81]int{
		5, 3, 4, 6, 7, 8, 9, 1, 2,
		6, 7, 2, 1, 9, 5, 3, 4, 8,
		1, 9, 8, 3, 4, 2, 5, 6, 7,
		8, 5, 9, 7, 6, 1, 4, 2, 3,
		4, 2, 6, 8, 5, 3, 7, 9, 1,
		7, 1, 3, 9, 2, 4, 8, 5, 6,
		9, 6, 1, 5, 3, 7, 2, 8, 4,
		2, 8, 7, 4, 1, 9, 6, 3, 5,
		3, 4, 5, 2, 8, 6, 1, 7, 9,
	}
}

func empty() [81]int {
	return [81]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
}

func UseVars() [][81]int {
	return [][81]int{empty(), sample1(), solutn1()}
}

func PrettyPrint(board [81]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", board[i*9+j])
			if j%3 == 2 {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
		if i%3 == 2 {
			fmt.Print("\n")
		}
	}
}

func ValidVars() {
	arr := game.ValidNumbers(sample1(), 2)
	fmt.Println(arr)
}
