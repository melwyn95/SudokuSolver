package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type sudokuResponse struct {
	Puzzle string `json:"puzzle"`
	Solved bool   `json:"solved"`
	Status string `json:"status"`
}

func SendResponse(w http.ResponseWriter, puzzle string, solved bool) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(sudokuResponse{
		Puzzle: puzzle,
		Solved: solved,
		Status: "success",
	})
}

func SendError(w http.ResponseWriter, message string) {
	http.Error(w, fmt.Errorf(message).Error(), http.StatusInternalServerError)
}
