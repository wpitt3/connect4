package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type GameResponse struct {
    BestMove int
    Result int
    Full bool
}

func game(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Access-Control-Allow-Origin")
  w.Header().Set("Content-Type", "application/json")

  if r.Method == "POST" {
    var board [7][6]int
    err := json.NewDecoder(r.Body).Decode(&board)
    if err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
    }
    bestMove := -1
    result := find4InBoard(board)
    full := result != 0 || boardIsFull(board)
    if (!full) {
      bestMove = findBestMove(board, 10000)
      var boardAfterMove = performMove(copyBoard(board), bestMove, 1)
      result = find4InBoard(boardAfterMove)
      full = result != 0 || boardIsFull(boardAfterMove)
    }

    response, err := json.Marshal(&GameResponse{BestMove: bestMove, Result: result, Full: full})
    if err != nil {
        fmt.Printf("Error: %s", err)
        return;
    }
    fmt.Fprintf(w, string(response))
  }

}

func main() {
	http.HandleFunc("/", game)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

