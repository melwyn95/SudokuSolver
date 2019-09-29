package tests

import (
	"fmt"
	"time"

	SudokuChecker "../sudokuchecker"
	SudokuSolver "../sudokusolver"
)

func RunTests() {
	easyPuzzle := "1.742..9.4.......2.5.7..1.39...5..6..4.8.6.2.3.5....1.2...13..6..3.9..7.6....84.5"
	start := time.Now()
	solvedEasyPuzzle, _ := SudokuSolver.Solve(easyPuzzle, false)
	end := time.Now()
	if SudokuChecker.CheckSolution(solvedEasyPuzzle) {
		fmt.Println("PASSED: Easy Puzzle", end.Sub(start))
	} else {
		fmt.Println("FAILED: Easy Puzzle", end.Sub(start))
	}

	mediumPuzzle := "8..4.6..7......4...1....65.5.9.3.78.....7.....48.2.1.3.52....9...1......3..9.2..5"
	start = time.Now()
	solvedMediumPuzzle, _ := SudokuSolver.Solve(mediumPuzzle, false)
	end = time.Now()
	if SudokuChecker.CheckSolution(solvedMediumPuzzle) {
		fmt.Println("PASSED: Medium Puzzle", end.Sub(start))
	} else {
		fmt.Println("FAILED: Medium Puzzle", end.Sub(start))
	}

	hardPuzzle := "8..........36......7..9.2...5...7.......457.....1...3...1....68..85...1..9....4.."
	start = time.Now()
	solvedHardPuzzle, _ := SudokuSolver.Solve(hardPuzzle, false)
	end = time.Now()
	if SudokuChecker.CheckSolution(solvedHardPuzzle) {
		fmt.Println("PASSED: Hard Puzzle", end.Sub(start))
	} else {
		fmt.Println("FAILED: Hard Puzzle", end.Sub(start))
	}

	invalidDuplicate := "8..........36......7..9.2...5...7.......457.....1...3...1....68..85...1..9....4.4"
	start = time.Now()
	solvedHardPuzzle, v1 := SudokuSolver.Solve(invalidDuplicate, false)
	end = time.Now()
	if !v1 {
		fmt.Println("PASSED: Invalid Puzzle", end.Sub(start))
	} else {
		fmt.Println("FAILED: Invalid Puzzle", end.Sub(start))
	}

	invalidIncomplete := "8..........36......7..9.2...5...7.......457.....1...3...1....68..85...1..9....4."
	start = time.Now()
	solvedHardPuzzle, v2 := SudokuSolver.Solve(invalidIncomplete, false)
	end = time.Now()
	if !v2 {
		fmt.Println("PASSED: Invalid Puzzle", end.Sub(start))
	} else {
		fmt.Println("FAILED: Invalid Puzzle", end.Sub(start))
	}
}
