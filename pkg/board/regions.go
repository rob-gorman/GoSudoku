package board

import "fmt"

// type region map[int]int // singular row/col/box { index: value }
type regionType [9]map[int]int
type regionFunc func(int) int
type RegionsByIndexes [BoardSize][3]map[int]int // {index: {regions including that index}}

var emptyBoard [BoardSize]int
var emptyRegionTemplate RegionsByIndexes // mapping of board index: [3]regionsForIndex

// empty regions as pre-compiled templates to build out actual instances
var rows regionType = buildFullRegion(emptyBoard, rowAlg)
var cols regionType = buildFullRegion(emptyBoard, colAlg)
var boxes regionType = buildFullRegion(emptyBoard, boxAlg)

// arrays for convenient iteration in helper functions
var regionsAlgs = [3]regionFunc{rowAlg, colAlg, boxAlg}
var regions = [3]regionType{rows, cols, boxes}

// Initializes empty Sudoku structure constants
// specifically `RegionsByIndexes` which is an array mapping each space (0-80)
// on the board to the index mappings of each region (row, column, box) it occupies
func Init() {
	initRegionsTemplate()
}

func initRegionsTemplate() {
	// this is constant and probably shouldn't be programmatically calculated
	// every time the application initializes, but it's fine
	for i := range emptyBoard {
		emptyRegionTemplate[i] = regionsForIndex(i)
	}
}

// returns an array of region mappings for a given index:
// {index: {{indexes in same row}, {same col}, {same box}}
func regionsForIndex(index int) (result [3]map[int]int) {
	for i := 0; i < len(regions); i++ {
		regionClass := regions[i]               // row, col, or box
		regionInstance := regionsAlgs[i](index) // specific row/col/box instance idx
		result[i] = regionClass[regionInstance] // assign region map to result idx
	}
	return result
}

// initializes mapping for each instance of a given region type
func buildFullRegion(board [81]int, fn regionFunc) (result regionType) {
	for regionNum := 0; regionNum < 9; regionNum++ {
		result[regionNum] = make(map[int]int)
	}

	// generate map keys as appropriate indicies of board structure
	for i := 0; i < BoardSize; i++ {
		regNum := fn(i)
		(result[regNum])[i] = board[i]
	}

	return result
}

// copies emptyRegionTemplate for mutation when initializing new GameBoard
func emptyRegionsByIndex() RegionsByIndexes {
	var template RegionsByIndexes

	for i := range emptyRegionTemplate {
		for region := range emptyRegionTemplate[i] {
			template[i][region] = make(map[int]int)

			for index, zeroVal := range emptyRegionTemplate[i][region] {
				template[i][region][index] = zeroVal
			}
		}
	}
	return template
}

// functions for programmatically determining to what region an index belongs
func rowAlg(index int) int {
	return index / 9
}

func colAlg(index int) int {
	return index % 9
}

func boxAlg(index int) int {
	row := rowAlg(index)
	col := colAlg(index)
	box := (row/3)*3 + (col / 3)
	return box
}

func prettyPrintIdxRegions(idx int, regions [3]map[int]int) string {
	var printStr = fmt.Sprintf("%d", idx)
	for _, region := range regions {
		printStr += fmt.Sprintf("\t%#v\n", region)
	}
	return printStr
}
