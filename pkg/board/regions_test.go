package board

import (
	"reflect"
	"sudoku/test_data"
	"testing"
)

func TestBuildFullRegion(t *testing.T) {
	tests := []struct {
		Name      string
		Board     [81]int
		Region    regionFunc
		SecondElt []int
		FourthElt []int
	}{
		{
			Name:      "empty board rows",
			Board:     test_data.Empty,
			Region:    rowAlg,
			SecondElt: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			FourthElt: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Name:      "sample1 board rows",
			Board:     test_data.Empty,
			Region:    rowAlg,
			SecondElt: []int{6, 0, 0, 1, 9, 5, 0, 0, 0},
			FourthElt: []int{8, 0, 0, 0, 6, 0, 0, 0, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			have := buildFullRegion(tc.Board, tc.Region)
			if !reflect.DeepEqual(have[1], tc.SecondElt) {
				t.Fatalf("regions not equal\nhave: %v\nwant: %v", have[1], tc.SecondElt)
			}
			if !reflect.DeepEqual(have[3], tc.FourthElt) {
				t.Fatalf("regions not equal\nhave: %v\nwant: %v", have[3], tc.FourthElt)
			}
		})
	}
}
