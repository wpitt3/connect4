package main

import (
    "math/rand"
    "time"
    "sort"
)

func abs(x int) int {
    if (x < 0) {
        return -x
    }
    return x
}

func copyBoard(board [7][6]int) [7][6]int {
    var newBoard [7][6]int
    for i:=0; i<7; i++ {
        for j:=0; j<6; j++ {
            newBoard[i][j] = board[i][j]
        }
    }
    return newBoard
}

func shuffle(list []int) []int {
    rand.Seed(time.Now().UTC().UnixNano())
    for i := len(list); i>0; i-- {
        ri := rand.Intn(i)
        temp := list[ri]
        list[ri] = list[i-1]
        list[i-1] = temp
    }
    return list
}

func sortChildren(children []*node){
    sort.Slice(children, func(i, j int) bool {
        return children[i].action < children[j].action
    })
}

func simulateGame(originalBoard [7][6]int, player int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    board := copyBoard(originalBoard)
    playerIndex := player
    result := find4InBoard(board)
    done := result != 0 || boardIsFull(board)
    for(!done) {
        moves := validMoves(board)
        move := moves[rand.Intn(len(moves))]
        board = performMove(board, move, playerIndex)
        playerIndex *= -1
        result = find4InBoard(board)
        done = result != 0 || boardIsFull(board)
    }
    return result
}
