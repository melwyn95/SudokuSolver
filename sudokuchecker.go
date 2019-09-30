package main

func checkPuzzle(grid *[9][9]uint8, rowIndex int, colIndex int, cell uint8) bool {
	row := GetRow(grid, rowIndex)
	row[colIndex] = 0

	col := GetCol(grid, colIndex)
	col[rowIndex] = 0

	box := GetBox(grid, rowIndex, colIndex)
	boxCenterX, boxCenterY := GetBoxCenterCoOrdinates(rowIndex, colIndex)
	boxIndex := (3 * (rowIndex - boxCenterX + 1)) + (colIndex - boxCenterY + 1)
	box[boxIndex] = 0

	return (ExistsArray(&row, cell) || ExistsArray(&col, cell) || ExistsArray(&box, cell))
}

func CheckSolution(solvedPuzzle string) bool {
	if len(solvedPuzzle) != 81 {
		return false
	}
	grid := ParsePuzzel(solvedPuzzle)
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			if checkPuzzle(&grid, rowIndex, colIndex, cell) {
				return false
			}
		}
	}
	return true
}

func ValidateSudoku(solvedPuzzle string) bool {
	nonEmptyCells := 0
	grid := ParsePuzzel(solvedPuzzle)
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			if cell != 0 {
				nonEmptyCells++
				if checkPuzzle(&grid, rowIndex, colIndex, cell) {
					return false
				}
			}
		}
	}
	return nonEmptyCells >= 17
}
