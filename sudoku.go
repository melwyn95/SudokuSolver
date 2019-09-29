package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	SudokuChecker "./sudokuchecker"
	SudokuResponse "./sudokuresponse"
	SudokuSolver "./sudokusolver"

	CacheUtils "./cacheutils"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	connection := CacheUtils.GetConnection()

	router.Path("/solve").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		puzzle := r.URL.Query().Get("puzzle")
		cacheKeySolve := CacheUtils.GenerateCacheKey(puzzle, "solve")
		cacheKeyCheck := CacheUtils.GenerateCacheKey(puzzle, "check")
		cachedSolution, ok := CacheUtils.GetKey(connection, cacheKeySolve)

		if !ok {
			solvedPuzzle, valid := SudokuSolver.Solve(puzzle, false)
			if !valid {
				http.Error(w, fmt.Errorf("Invalid Sudoku").Error(), http.StatusInternalServerError)
				return
			}
			solved := SudokuChecker.CheckSolution(solvedPuzzle)

			CacheUtils.SetKey(connection, cacheKeySolve, solvedPuzzle)
			CacheUtils.SetKey(connection, cacheKeyCheck, strconv.FormatBool(solved))

			SudokuResponse.SendResponse(w, solvedPuzzle, solved)
			return
		}

		cachedCheck, _ := CacheUtils.GetKey(connection, cacheKeyCheck)
		solved, _ := strconv.ParseBool(cachedCheck)
		SudokuResponse.SendResponse(w, cachedSolution, solved)
	})

	router.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		puzzle := r.URL.Query().Get("puzzle")
		cacheKeyCheck := CacheUtils.GenerateCacheKey(puzzle, "check")
		cachedCheck, ok := CacheUtils.GetKey(connection, cacheKeyCheck)
		if !ok {
			solved := SudokuChecker.CheckSolution(puzzle)
			if solved {
				CacheUtils.SetKey(connection, cacheKeyCheck, "true")
				SudokuResponse.SendResponse(w, puzzle, solved)
				return
			}
			CacheUtils.SetKey(connection, cacheKeyCheck, "false")
			SudokuResponse.SendError(w, "Invalid Sudoku")
			return
		}
		solved, _ := strconv.ParseBool(cachedCheck)
		SudokuResponse.SendResponse(w, puzzle, solved)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
