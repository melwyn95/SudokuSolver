package sudokuchecker

import (
	ArrayUtils "../arrayutils"
	SudokuIO "../sudokuio"
	SudokuUtils "../sudokuutils"
)

func checkPuzzle(grid *[9][9]uint8, rowIndex int, colIndex int, cell uint8) bool {
	row := SudokuUtils.GetRow(grid, rowIndex)
	row[colIndex] = 0

	col := SudokuUtils.GetCol(grid, colIndex)
	col[rowIndex] = 0

	box := SudokuUtils.GetBox(grid, rowIndex, colIndex)
	boxCenterX, boxCenterY := SudokuUtils.GetBoxCenterCoOrdinates(rowIndex, colIndex)
	boxIndex := (3 * (rowIndex - boxCenterX + 1)) + (colIndex - boxCenterY + 1)
	box[boxIndex] = 0

	return (ArrayUtils.Exists(&row, cell) || ArrayUtils.Exists(&col, cell) || ArrayUtils.Exists(&box, cell))
}

func CheckSolution(solvedPuzzle string) bool {
	if len(solvedPuzzle) != 81 {
		return false
	}
	grid := SudokuIO.ParsePuzzel(solvedPuzzle)
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
	grid := SudokuIO.ParsePuzzel(solvedPuzzle)
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
