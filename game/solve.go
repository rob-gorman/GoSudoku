package game

func Solve(board [81]int) (solution [81]int, solved bool) {
	if !ValidateBoard(board) {
		return *new([81]int), false
	}
	return populateBoard(board)
}

func populateBoard(board [81]int) ([81]int, bool) {
	newBoard := [81]int{}
	return newBoard, false
}
