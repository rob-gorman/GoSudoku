sudoku
 |
 |- Server -> board to app
 |
 |- Regions
 |- App


Server
App
Fill Square
Validate

# Handler
- Accept HTTP request
- decode body into boardinput (map[int]int)
- pass boardinput to solver

# Regions
- Init build regions by index tmplate ([81][3]map[int]int // {index: value}, value init 0)
  - initRegionTemplate() {}
  - iterate over empty board idxs
  - result[idx] = regionsForIndex(idx)
- Use empty region template as existing structure to build out index regions for board
- ValidValues = []int of possible values for a single space -> what is exported to app

# Solve
- validate input board (?)
- find first open square
- solve board (board, next index)
  - if next index == len board return board // solved
  - validMoves = []int
  - if len == 0 return err
  - PopulateSquare: iterate validMoves => board
  - return solve board (nextBoard, index+1)
  - return board arr, solved bool

