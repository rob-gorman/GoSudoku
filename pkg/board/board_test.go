package board

import (
	"os"
	"reflect"
	"sudoku/test_data"
	"testing"
)

func TestMain(m *testing.M) {
	Init()
	os.Exit(m.Run())
}

func TestInsertRegionValues(t *testing.T) {
	tests := []struct {
		Name     string
		Input    [81]int
		Elt0 int
		Elt61 int
	}{
		{
			Name:     "empty board",
			Input:    test_data.Empty,
			Elt0: 0,
			Elt61: 0,
		},
		{
			Name:     "standard board",
			Input:    test_data.Sample1,
			Elt0: 5,
			Elt61: 2,
		},
	}

	for _, tc := range tests {
		gb := New()
		gb.Board = tc.Input
		for i := range gb.Board {
			gb.indexRegionsValues(i)
		}

		elt1Regions := gb.Regions[0]
		elt1inRow := elt1Regions[0][0]
		elt2Regions := gb.Regions[61]
		elt2inCol := elt2Regions[1][61]

		t.Run(tc.Name, func(t *testing.T) {
			if !reflect.DeepEqual(tc.Input, gb.Board) {
				t.Fatalf("board slice not initialized correctly\n%s", test_data.PrettyPrint(gb.Board))
			}
			if elt1inRow != tc.Elt0 {
				printstr := ""
				for i := range gb.Regions {
					printstr += prettyPrintIdxRegions(i, gb.Regions[i])
				}
				t.Fatalf("rows inconsistent. have: %d; want: %d\n%s", elt1inRow, tc.Elt0, printstr)
			}
			if elt2inCol != tc.Elt61 {
				printstr := ""
				for i := range gb.Regions {
					printstr += prettyPrintIdxRegions(i, gb.Regions[i])
				}
				t.Fatalf("regions inconsistent. have: %d; want: %d\nregions: %s", elt2inCol, tc.Elt61, printstr)
			}
		})
	}
}

// func TestGBFromBoard(t *testing.T) {
// 	tests := []struct {
// 		Name     string
// 		Input    [81]int
// 		Elt0 int
// 		Elt61 int
// 	}{
// 		{
// 			Name:     "empty board",
// 			Input:    test_data.Empty,
// 			Elt0: 0,
// 			Elt61: 0,
// 		},
// 		{
// 			Name:     "standard board",
// 			Input:    test_data.Sample1,
// 			Elt0: 5,
// 			Elt61: 2,
// 		},
// 	}

// 	for _, tc := range tests {
// 		have := gameBoardFromBoard(&tc.Input)
// 		elt1Regions := have.Regions[0]
// 		elt1inRow := elt1Regions[0][0]
// 		elt2Regions := have.Regions[61]
// 		elt2inCol := elt2Regions[1][61]

// 		t.Run(tc.Name, func(t *testing.T) {
// 			if !reflect.DeepEqual(tc.Input, have.Board) {
// 				t.Fatalf("board slice not initialized correctly\n%s", test_data.PrettyPrint(have.Board))
// 			}
// 			if elt1inRow != tc.Elt0 {
// 				t.Fatalf("rows inconsistent. have: %d; want: %d", elt1inRow, tc.Elt0)
// 			}
// 			if elt2inCol != tc.Elt61 {
// 				printstr := ""
// 				for i := range have.Regions {
// 					printstr += prettyPrintIdxRegions(i, have.Regions[i])
// 				}
// 				t.Fatalf("regions inconsistent. have: %d; want: %d\nregions: %s", elt2inCol, tc.Elt61, printstr)
// 			}
// 		})
// 	}
// }

func TestValidate(t *testing.T) {
	tests := []struct {
		Name  string
		Input [81]int
		Want  error
	}{
		{
			Name:  "valid board",
			Input: test_data.Sample1,
			Want:  nil,
		},
		{
			Name:  "empty board",
			Input: test_data.Empty,
			Want:  nil,
		},
		{
			Name:  "invalid-repeat first two elts",
			Input: test_data.InvalidBoard1,
			Want:  ErrInvalidPuzzle,
		},
		{
			Name:  "invalid-opposite box corners",
			Input: test_data.InvalidBoard2,
			Want:  ErrInvalidPuzzle,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			gb := gameBoardFromBoard(&tc.Input)
			have := gb.validate()
			if have != tc.Want {
				t.Fatalf("incorrect output: %v", have)
			}
		})
	}
}
