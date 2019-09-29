package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	SudokuChecker "./sudokuchecker"
	SudokuResponse "./sudokuresponse"
	SudokuSolver "./sudokusolver"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Path("/solve").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		puzzle := r.URL.Query().Get("puzzle")
		solvedPuzzle, valid := SudokuSolver.Solve(puzzle, false)
		if !valid {
			http.Error(w, fmt.Errorf("Invalid Sudoku").Error(), http.StatusInternalServerError)
			return
		}
		solved := SudokuChecker.CheckSolution(solvedPuzzle)

		SudokuResponse.SendResponse(w, solvedPuzzle, solved)
	})

	router.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		puzzle := r.URL.Query().Get("puzzle")
		solved := SudokuChecker.CheckSolution(puzzle)
		if solved {
			SudokuResponse.SendResponse(w, puzzle, solved)
		} else {
			http.Error(w, fmt.Errorf("Invalid Sudoku").Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
