package SudokuChecker

import (
	"../ArrayUtils"
	"../SudokuIO"
	"../SudokuUtils"
)

func Check(solvedPuzzle string) bool {
	grid := SudokuIO.ParsePuzzel(solvedPuzzle)
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			row := SudokuUtils.GetRow(&grid, rowIndex)
			row[colIndex] = 0

			col := SudokuUtils.GetCol(&grid, colIndex)
			col[rowIndex] = 0

			box := SudokuUtils.GetBox(&grid, rowIndex, colIndex)
			boxCenterX, boxCenterY := SudokuUtils.GetBoxCenterCoOrdinates(rowIndex, colIndex)
			boxIndex := (3 * (rowIndex - boxCenterX + 1)) + (colIndex - boxCenterY + 1)
			box[boxIndex] = 0

			if ArrayUtils.Exists(&row, cell) || ArrayUtils.Exists(&col, cell) || ArrayUtils.Exists(&box, cell) {
				return false
			}
		}
	}
	return true
}
