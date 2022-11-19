package api

import (
	"encoding/json"
	"fmt"
	"sudoku/pkg/board"
)

type Input struct {
	In map[int]int
}

// Errors
var ErrInvalidIndex = fmt.Errorf("invalid index: out of range")

func (in Input) ToBoard() (bd *[board.BoardSize]int, err error) {
	for k, v := range in.In {
		if k >= board.BoardSize {
			return nil, ErrInvalidIndex
		}
		bd[k] = v
	}
	return bd, nil
}

func FromJSON(req []byte) (*[board.BoardSize]int, error) {
	// unnecessary function wrapping of ToBoard?
	submission := &Input{
		In: make(map[int]int),
	}
	err := json.Unmarshal(req, submission.In)
	if err != nil {
		return nil, err
	}

	return submission.ToBoard()
}
