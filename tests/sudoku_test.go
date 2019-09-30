package tests

import (
	"os"
	"testing"

	SudokuChecker "../sudokuchecker"
	SudokuSolver "../sudokusolver"
)

func TestSudoku(t *testing.T) {
	t.Run("Easy Puzzle", func(t *testing.T) {
		easyPuzzle := "1.742..9.4.......2.5.7..1.39...5..6..4.8.6.2.3.5....1.2...13..6..3.9..7.6....84.5"
		solvedEasyPuzzle, _ := SudokuSolver.Solve(easyPuzzle, false)
		if !SudokuChecker.CheckSolution(solvedEasyPuzzle) {
			t.Errorf("FAILED: Easy Puzzle")
		}
	})

	t.Run("Medium Puzzle", func(t *testing.T) {
		mediumPuzzle := "8..4.6..7......4...1....65.5.9.3.78.....7.....48.2.1.3.52....9...1......3..9.2..5"
		solvedMediumPuzzle, _ := SudokuSolver.Solve(mediumPuzzle, false)
		if !SudokuChecker.CheckSolution(solvedMediumPuzzle) {
			t.Errorf("FAILED: Medium Puzzle")
		}
	})

	t.Run("Hard Puzzle", func(t *testing.T) {
		hardPuzzle := "8..........36......7..9.2...5...7.......457.....1...3...1....68..85...1..9....4.."
		solvedHardPuzzle, _ := SudokuSolver.Solve(hardPuzzle, false)
		if !SudokuChecker.CheckSolution(solvedHardPuzzle) {
			t.Errorf("FAILED: Hard Puzzle")
		}
	})

	t.Run("Invalid Puzzle 1", func(t *testing.T) {
		invalidDuplicate := "8..........36......7..9.2...5...7.......457.....1...3...1....68..85...1..9....4.4"
		_, v1 := SudokuSolver.Solve(invalidDuplicate, false)
		if v1 {
			t.Errorf("FAILED: Invalid Puzzle 1")
		}
	})

	t.Run("Invalid Puzzle 2", func(t *testing.T) {
		invalidIncomplete := "8..........36......7..9.2...5...7.......457.....1...3...1....68..85...1..9....4."
		_, v2 := SudokuSolver.Solve(invalidIncomplete, false)
		if v2 {
			t.Errorf("FAILED: Invalid Puzzle 2")
		}
	})

}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
