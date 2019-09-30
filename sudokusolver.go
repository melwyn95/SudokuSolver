package main

func narrowPossibilies(grid *[9][9]uint8, rowIndex int, colIndex int) []uint8 {
	var narrow []uint8
	if (*grid)[rowIndex][colIndex] > 0 {
		return narrow
	}
	row, col, box := GetRow(grid, rowIndex), GetCol(grid, colIndex), GetBox(grid, rowIndex, colIndex)
	for _, possibility := range GetAllPossibiliesForCell() {
		if !(ExistsArray(&row, possibility) || ExistsArray(&col, possibility) || ExistsArray(&box, possibility)) {
			narrow = append(narrow, possibility)
		}
	}
	return narrow
}

func generatePossibilitiesForBlankCells(grid *[9][9]uint8) [9][9][]uint8 {
	var possibilitiesGrid [9][9][]uint8
	for rowIndex, row := range grid {
		for colIndex, _ := range row {
			possibilitiesGrid[rowIndex][colIndex] = narrowPossibilies(grid, rowIndex, colIndex)
		}
	}
	return possibilitiesGrid
}

func eliminatePossiblities(rowIndex int, colIndex int, grid *[9][9]uint8, possibilitiesGrid *[9][9][]uint8) *[9][9][]uint8 {
	cellValue := (*grid)[rowIndex][colIndex]
	for index := 0; index < 9; index++ {
		sliceIndex := FindIndexSlice((*possibilitiesGrid)[rowIndex][index], cellValue)
		if sliceIndex >= 0 {
			(*possibilitiesGrid)[rowIndex][index] = RemoveFromSlice((*possibilitiesGrid)[rowIndex][index], sliceIndex)
		}

		sliceIndex = FindIndexSlice((*possibilitiesGrid)[index][colIndex], cellValue)
		if sliceIndex >= 0 {
			(*possibilitiesGrid)[index][colIndex] = RemoveFromSlice((*possibilitiesGrid)[index][colIndex], sliceIndex)
		}
	}
	centerX, centerY := GetBoxCenterCoOrdinates(rowIndex, colIndex)
	for rowOffset := -1; rowOffset < 2; rowOffset++ {
		for colOffset := -1; colOffset < 2; colOffset++ {
			boxX, boxY := centerX+rowOffset, centerY+colOffset
			sliceIndex := FindIndexSlice((*possibilitiesGrid)[boxX][boxY], cellValue)
			if sliceIndex >= 0 {
				(*possibilitiesGrid)[boxX][boxY] = RemoveFromSlice((*possibilitiesGrid)[boxX][boxY], sliceIndex)
			}
		}
	}
	return possibilitiesGrid
}

func eliminateObviousPossibilitesFromPosssibilitiesGrid(grid *[9][9]uint8, possibilitiesGrid *[9][9][]uint8) ([9][9]uint8, [9][9][]uint8) {
	for true {
		possibilitiesEliminated := false
		for rowIndex, row := range possibilitiesGrid {
			for colIndex, cell := range row {
				if len(cell) == 1 {
					(*grid)[rowIndex][colIndex] = cell[0]
					possibilitiesGrid = eliminatePossiblities(rowIndex, colIndex, grid, possibilitiesGrid)
					possibilitiesEliminated = true
				}
			}
		}
		if !possibilitiesEliminated {
			break
		}
	}
	return *grid, *possibilitiesGrid
}

func copyPossibilitiesGrid(possibilitiesGrid *[9][9][]uint8) *[9][9][]uint8 {
	var newPossibilitiesGrid [9][9][]uint8
	for rowIndex, row := range possibilitiesGrid {
		for colIndex, cellSlice := range row {
			newSlice := make([]uint8, len(cellSlice))
			copy(newSlice, cellSlice)
			newPossibilitiesGrid[rowIndex][colIndex] = newSlice
		}
	}
	return &newPossibilitiesGrid
}

func solve(grid [9][9]uint8, possibilitiesGrid *[9][9][]uint8) ([9][9]uint8, bool) {
	emptyX, emptyY := FindBlankCell(&grid)
	if emptyX >= 0 && emptyY >= 0 {
		for _, possibility := range (*possibilitiesGrid)[emptyX][emptyY] {
			grid[emptyX][emptyY] = possibility
			newPossibilitiesGrid := eliminatePossiblities(emptyX, emptyY, &grid, copyPossibilitiesGrid(possibilitiesGrid))
			newGrid, solved := solve(grid, newPossibilitiesGrid)
			if solved {
				return newGrid, true
			}
		}
	} else {
		return grid, true
	}
	return grid, false
}

func Solve(puzzle string, verbose bool) (string, bool) {
	if len(puzzle) != 81 {
		return puzzle, false
	}
	grid := ParsePuzzel(puzzle)
	if !ValidateSudoku(puzzle) {
		return puzzle, false
	}

	if verbose {
		PrintPuzzle(&grid)
	}
	possibilitiesGrid := generatePossibilitiesForBlankCells(&grid)
	grid, possibilitiesGrid = eliminateObviousPossibilitesFromPosssibilitiesGrid(&grid, &possibilitiesGrid)
	grid, _ = solve(grid, &possibilitiesGrid)
	if verbose {
		PrintPuzzle(&grid)
	}
	return StringifyPuzzle(&grid), true
}
