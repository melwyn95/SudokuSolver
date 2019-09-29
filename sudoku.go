package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./SudokuChecker"
	"./SudokuResponse"
	"./SudokuSolver"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Path("/solve").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		puzzle := r.URL.Query().Get("puzzle")
		solvedPuzzle := SudokuSolver.Solve(puzzle, false)
		solved := SudokuChecker.Check(solvedPuzzle)

		SudokuResponse.SendResponse(w, solvedPuzzle, solved)
	})

	router.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		puzzle := r.URL.Query().Get("puzzle")
		solved := SudokuChecker.Check(puzzle)

		SudokuResponse.SendResponse(w, puzzle, solved)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
