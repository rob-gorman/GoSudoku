package api

import (
	"encoding/json"
	"sudoku/pkg/board"
)

type SolutionResponse struct {
	Solution [board.BoardSize]int `json:"solution"`
}

func (s *SolutionResponse) ToJSON() (res []byte, err error) {
	return json.Marshal(s)
}