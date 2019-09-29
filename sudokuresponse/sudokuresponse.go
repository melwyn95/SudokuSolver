package sudokuresponse

import (
	"encoding/json"
	"net/http"
)

type sudokuResponse struct {
	Puzzle string `json:"puzzle"`
	Solved bool   `json:"solved"`
	Status string `json:"status"`
}

func SendResponse(w http.ResponseWriter, puzzle string, solved bool) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sudokuResponse{
		Puzzle: puzzle,
		Solved: solved,
		Status: "success",
	})
}
