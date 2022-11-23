package board

import "fmt"

const BoardSize = 81

var ErrInvalidPuzzle = fmt.Errorf("invalid puzzle submission")

type GameBoard struct {
	Board   [81]int
	Regions RegionsByIndexes
}

// generates new GameBoard from flat input array
// returns an error if board violates sudoku rules
// however it assumes valid inputs (1-9) within bounds
func MakeBoard(input *[81]int) (board *GameBoard, err error) {
	board = gameBoardFromBoard(input)
	err = board.validate()
	return board, err
}

// intializes a new empty board
func New() *GameBoard {
	return &GameBoard{
		Board:   [81]int{},
		Regions: emptyRegionsByIndex(),
	}
}

// parses a flat array of 81-space board into GameBoard struct
func gameBoardFromBoard(input *[81]int) *GameBoard {
	gb := New()
	for i, v := range input {
		gb.Board[i] = v
		gb.insertRegionsValue(i)
	}
	return gb
}

func (b *GameBoard) LegalMoves(idx int) []int {
	result := make([]int, 0, 9)
	existingValues := b.indexRegionsValues(idx)
	for i := 1; i <= 9; i++ {
		if !existingValues[i] {
			result = append(result, i)
		}
	}
	return result
}

func (b GameBoard) FillValue(idx, val int) GameBoard {
	b.Board[idx] = val
	b.indexRegionsValues(idx)
	return b
}

func (b *GameBoard) insertRegionsValue(idx int) {
	value := b.Board[idx]
	indexRegs := b.Regions[idx]
	for _, region := range indexRegs {
		region[idx] = value
	}
}

func (b *GameBoard) indexRegionsValues(idx int) map[int]bool {
	result := make(map[int]bool)
	for _, region := range b.Regions[idx] {
		for _, v := range region {
			result[v] = true
		}
		delete(region, 0)
	}
	return result
}

func (b *GameBoard) validate() error {
	for index := range b.Regions {
		for _, region := range b.Regions[index] {
			regionVals := make(map[int]bool)
			for _, val := range region {
				if val == 0 {
					continue
				}
				if regionVals[val] {
					return ErrInvalidPuzzle
				}
				regionVals[val] = true
			}
		}
	}

	return nil
}
