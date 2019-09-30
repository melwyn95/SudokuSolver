package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter().StrictSlash(true)
	connection := GetConnection()

	router.Path("/solve").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			return
		}

		puzzle := r.URL.Query().Get("puzzle")
		cacheKeySolve := GenerateCacheKey(puzzle, "solve")
		cacheKeyCheck := GenerateCacheKey(puzzle, "check")
		cachedSolution, ok := GetKey(connection, cacheKeySolve)

		if !ok {
			solvedPuzzle, valid := Solve(puzzle, false)
			if !valid {
				http.Error(w, fmt.Errorf("Invalid Sudoku").Error(), http.StatusInternalServerError)
				return
			}
			solved := CheckSolution(solvedPuzzle)

			SetKey(connection, cacheKeySolve, solvedPuzzle)
			SetKey(connection, cacheKeyCheck, strconv.FormatBool(solved))

			SendResponse(w, solvedPuzzle, solved)
			return
		}

		cachedCheck, _ := GetKey(connection, cacheKeyCheck)
		solved, _ := strconv.ParseBool(cachedCheck)
		SendResponse(w, cachedSolution, solved)
	}).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			return
		}

		puzzle := r.URL.Query().Get("puzzle")
		cacheKeyCheck := GenerateCacheKey(puzzle, "check")
		cachedCheck, ok := GetKey(connection, cacheKeyCheck)
		if !ok {
			solved := CheckSolution(puzzle)
			if solved {
				SetKey(connection, cacheKeyCheck, "true")
				SendResponse(w, puzzle, solved)
				return
			}
			SetKey(connection, cacheKeyCheck, "false")
			SendError(w, "Invalid Sudoku")
			return
		}
		solved, _ := strconv.ParseBool(cachedCheck)
		SendResponse(w, puzzle, solved)
	}).Methods(http.MethodGet, http.MethodOptions)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
