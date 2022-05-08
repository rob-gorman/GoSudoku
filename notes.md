## Sudoku Backend

# Flow
- incoming JSON as [81]int
  - `0` as empty space
  - _validate_ intial board

- fill board
  - while valid returns true, return recursively filled next square 1-9
    - maybe optimal to get available squares here from rows
  - otherwise, valid returns false
  - if board can't be filled?

- return board as JSON
  - figure this out

# Tricky Parts
**VALIDATION**
- _Representation of Regions:_
  - Cols, Rows, Boxes
  - Each arrays of `maps[int]int` => { index: value }

- _Validating active square vs Entire Board_

- _Validating active square_
  - Retrieve map from each subarray (3x 9-key map elements) containing the index of space being filled
  - Return true/false value already exists in any subarray

- _Validating entire board_
  - validate that no element occurs in any map element twice
  - loop through subarrays, passing each to `validateRegion` function

**FILL SQUARE**
- index of first `0` in array
- retrieve regions containing index (3 maps)
- create `[]int` of valid numbers
- iterate over value slice to iteratively fill square
- if valid numbers is empty ?? ...return _, false
- board filled return board, true

```go
// if len(validNumbers(index)) == 0 { return [81]int, false }
for i, v := range validNumbers {
  newboard := board.insertValue(ind, v)
  if newboard.isFull() { return newboard, true }

  return { newboard.fillSquare(getNextSpace()) }
}
```
