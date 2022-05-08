package main

import (
	"fmt"
	"sudoku/testing"
)

func main() {
	toPrint := testing.UseVars()
	for i := 0; i < len(toPrint); i++ {
		testing.PrettyPrint(toPrint[i])
		fmt.Print("\n")
	}

	testing.ValidVars()
}
