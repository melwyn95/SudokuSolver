package sudokuutils

func FindBlankCell(grid *[9][9]uint8) (int, int) {
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			if cell == 0 {
				return rowIndex, colIndex
			}
		}
	}
	return -1, -1
}

func GetAllPossibiliesForCell() [9]uint8 {
	var possibilies [9]uint8
	var n uint8 = 1
	for ; n <= 9; n++ {
		possibilies[n-1] = n
	}
	return possibilies
}

func GetBoxCenterCoOrdinates(rowIndex, colIndex int) (int, int) {
	return 3*(rowIndex/3) + 1, 3*(colIndex/3) + 1
}

func GetRow(grid *[9][9]uint8, rowIndex int) [9]uint8 {
	return (*grid)[rowIndex]
}

func GetCol(grid *[9][9]uint8, colIndex int) [9]uint8 {
	var col [9]uint8
	for rowIndex := 0; rowIndex < 9; rowIndex++ {
		col[rowIndex] = (*grid)[rowIndex][colIndex]
	}
	return col
}

func GetBox(grid *[9][9]uint8, rowIndex int, colIndex int) [9]uint8 {
	var box [9]uint8
	boxIndex := 0
	centerX, centerY := GetBoxCenterCoOrdinates(rowIndex, colIndex)
	for rowOffset := -1; rowOffset < 2; rowOffset++ {
		for colOffset := -1; colOffset < 2; colOffset++ {
			box[boxIndex] = (*grid)[centerX+rowOffset][centerY+colOffset]
			boxIndex++
		}
	}
	return box
}
