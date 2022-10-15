package game

func Solve(board [81]int) (solution *[81]int, solved bool) {
	if !ValidateBoard(board) {
		return nil, false
	}
	return solveBoard(board)
}

func solveBoard(board [81]int) (*[81]int, bool) {
	openSquare, full := nextSpaceIndex(board)
	if full {
		return &board, true
	}

	moves := validNumbers(board, *openSquare)
	if len(moves) == 0 {
		return nil, false
	}

	// concurrently?
	for i := range moves {
		newBoard := fillSquare(board, *openSquare, moves[i])
		if solution, ok := solveBoard(newBoard); ok {
			return solution, ok
		}
	}

	return nil, false
}

// finds next open space
// if none, the board is full
func nextSpaceIndex(board [81]int) (*int, bool) {
	for i, v := range board {
		if v == 0 {
			return &i, true
		}
	}
	return nil, false
}

// returns the currently valid numbers for a given index
func validNumbers(board [81]int, index int) (result []int) {
	targetRegions := regionsIncludingIndex(index)
	fillIndexRegions(board, targetRegions)
	filledValues := valuesInRegions(targetRegions)

	for i := 1; i <= 9; i++ {
		if !filledValues[i] {
			result = append(result, i)
		}
	}

	return result
}

func fillSquare(board [81]int, ind int, value int) (result [81]int) {
	board[ind] = value
	return board
}
