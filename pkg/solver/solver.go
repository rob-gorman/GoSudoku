package solver

import (
	"sudoku/pkg/board"
)

func Solve(input *[board.BoardSize]int) (solution [board.BoardSize]int, err error) {
	gameBoard, err := board.New(input)
	if err != nil {
		// log
		return solution, err
	}

	solved, err := populateBoard(*gameBoard, 0)
	if err != nil {
		// log
		return solution, err
	}
	return solved.Board, err
}

func populateBoard(bd board.GameBoard, idx int) (next board.GameBoard, err error) {
	if idx > 80 {
		return bd, err
	}
	if bd.Board[idx] != 0 { // space already filled
		return populateBoard(bd, idx+1)
	}

	moves := bd.LegalMoves(idx)
	for _, val := range moves {
		next = bd.FillValue(idx, val)
		solution, err := populateBoard(next, idx+1)
		if err != nil {
			continue
		}
		return solution, nil
	}

	// no valid moves
	return next, err
}
