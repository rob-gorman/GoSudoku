package api

import (
	"sudoku/pkg/solver"
)

func SolveBoard(req []byte) (res []byte, err error) {
	bd, err := FromJSON(req)
	if err != nil {
		return nil, err
	}

	solution, err := solver.Solve(bd)
	if err != nil {
		return nil, err
	}

	out := SolutionResponse{solution}
	return out.ToJSON()
}
