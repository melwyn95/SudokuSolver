package sudokuio

import (
	"fmt"
	"strconv"
)

func ParsePuzzel(puzzle string) [9][9]uint8 {
	var grid [9][9]uint8
	for index, chr := range puzzle {
		var v, e = strconv.ParseUint(string(chr), 10, 8)
		if e != nil {
			grid[index/9][index%9] = 0
		} else {
			grid[index/9][index%9] = uint8(v)
		}
	}
	return grid
}

func StringifyPuzzle(grid *[9][9]uint8) string {
	puzzle := ""
	for _, row := range grid {
		for _, cell := range row {
			puzzle += strconv.Itoa(int(cell))
		}
	}
	return puzzle
}

func PrintPuzzle(grid *[9][9]uint8) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}
