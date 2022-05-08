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

# Regions
- Build each subregion as { boardIndex(0-80) => value(1-9) } (for next step)
- Use empty region template as existing structure to build out index regions
- ValidValues = []int of possible values for a single space -> what is exported to app

# Solve
- validate board (?)
- PopulateBoard: iterate validNumbers
- return board arr, solved bool