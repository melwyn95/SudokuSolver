package main

import (
	"fmt"
	"time"

	"./SudokuChecker"
	"./SudokuSolver"
	"./Tests"
)

func main() {
	Tests.RunTests()

	veryHardPuzzle := "8..........36......7..9.2...5...7.......457.....1...3...1....68..85...1..9....4.."
	start := time.Now()

	solvedPuzzle := SudokuSolver.Solve(veryHardPuzzle, false)

	end := time.Now()
	fmt.Println(end.Sub(start))

	solved := SudokuChecker.Check(solvedPuzzle)
	if solved {
		fmt.Println("!! Solved !!")
	} else {
		fmt.Println("XX Wrong Answer XX")
	}
}
