package main

import (
	"testing"
	"fmt"
)

func shite(t *testing.T) {
    fmt.Printf("Go says %v\n", 0)
}

func Test_findGridColour(t *testing.T) {
    equal(t, 0, findGridColour(0, 0))
    equal(t, 1, findGridColour(1, 0))
    equal(t, 4, findGridColour(0, 1))
    equal(t, 3, findGridColour(0, 2))
    equal(t, 2, findGridColour(0, 3))
    equal(t, 1, findGridColour(0, 4))
}

func Test_findGridColumn(t *testing.T) {
    equal(t, 0, findGridColumn(0, 0))
    equal(t, 1, findGridColumn(0, 1)) // blue on row 2
    equal(t, 2, findGridColumn(0, 2)) // blue on row 3
    equal(t, 0, findGridColumn(4, 1))
    equal(t, 0, findGridColumn(1, 4))
}

func Test_findValidColumnColours_allPossibleWhenEmpty(t *testing.T) {
    state := newState()

    result := findValidColumnColours(state)

    equal(t, 25, int8(len(result)))
}

func Test_findValidColumnColours_rowNotPossibleWhenFull(t *testing.T) {
    state := newState()

    state.boards[0].rows[0][1] = 1
    state.boards[0].rows[1][1] = 2

    result := findValidColumnColours(state)

    equal(t, 15, int8(len(result)))
}

func Test_findValidColumnColours_onlyMatchingColourInRowIsPossible(t *testing.T) {
    state := newState()

    state.boards[0].rows[1][0] = 3 // colour
    state.boards[0].rows[1][1] = 1

    result := findValidColumnColours(state)

    equal(t, 21, int8(len(result)))
    equalBool(t, true, result[1*5+3])
}

func Test_findValidColumnColours_onlyEmptyGridIsPossible(t *testing.T) {
    state := newState()

    state.boards[0].grid[2][findGridColumn(0,2)] = true
    state.boards[0].grid[2][findGridColumn(1,2)] = true
    state.boards[0].grid[2][findGridColumn(3,2)] = true

    result := findValidColumnColours(state)

    equal(t, 22, int8(len(result)))
    equalBool(t, true, result[2*5+2])
    equalBool(t, true, result[2*5+4])
}

func Test_findValidMoves_emptyBoardWithSinglePile(t *testing.T) {
    state := newState()
    state = appendPiles(state, 1)
    state.piles[0][1] = 2

    result := findValidMoves(state)

    equal(t, 6, int8(len(result)))
    for i:=0; i<6; i++ {
        equal(t, 0, result[0].pilesIndex)
        equal(t, 1, result[0].colour)
    }
}

func Test_findValidMoves_oneRowFullWithSinglePile(t *testing.T) {
    state := newState()
    state = appendPiles(state, 1)
    state.piles[0][1] = 2
    state.boards[0].rows[0][1] = 1

    result := findValidMoves(state)

    equal(t, 5, int8(len(result)))
}

func Test_isRoundOver_emptyPile(t *testing.T) {
    state := newState()
    state = appendPiles(state, 1)

    result := isRoundOver(state)

    equalBool(t, true, result)
}

func Test_isRoundOver_notOver(t *testing.T) {
    state := newState()
    state = appendPiles(state, 1)
    state.piles[0][1] = 2

    result := isRoundOver(state)

    equalBool(t, false, result)
}

func Test_isGameOver_notOver(t *testing.T) {
     state := newState()

     result := isGameOver(state)

     equalBool(t, false, result)
}

func Test_isGameOver_isOver(t *testing.T) {
     state := newState()
     for i:=0; i<5; i++ {
        state.boards[1].grid[i][1] = true
     }

     result := isGameOver(state)

     equalBool(t, true, result)
 }

func Test_countAdjacent_line(t *testing.T) {
    board := Board{}
    board.grid[1][0] = true
    board.grid[1][1] = true
    board.grid[1][2] = true
    board.grid[1][4] = true

    result := countAdjacent(board, 1, 1)

    equal(t, 3, int8(result))
}

func Test_countAdjacent_cross(t *testing.T) {
    board := Board{}

    board.grid[1][0] = true
    board.grid[1][1] = true
    board.grid[1][2] = true
    board.grid[0][1] = true

    result := countAdjacent(board, 1, 1)

    equal(t, 5, int8(result))
}

func Test_endRound_cross(t *testing.T) {
    state := newState()

    state.boards[0].rows[1][0]=0 // colour 0
    state.boards[0].rows[1][1]=2 // index 1
    state.boards[0].rows[4][1]=5 // on its own
    state.boards[0].rows[3][1]=3 // not full

    state.boards[0].grid[1][0] = true
    state.boards[0].grid[1][2] = true

    state = endRound(state)

    equal(t, 4, int8(state.boards[0].score))
    equal(t, 0, int8(state.boards[0].rows[1][1]))
    equal(t, 0, int8(state.boards[0].rows[4][1]))
    equal(t, 3, int8(state.boards[0].rows[3][1]))
}

func Test_endRound_discard_overfull(t *testing.T) {
    state := newState()

    state.boards[0].score = 100
    state.boards[0].discard = 10

    state = endRound(state)

    equal(t, 86, int8(state.boards[0].score))
}

func Test_endRound_discard_negativeScore(t *testing.T) {
    state := newState()

    state.boards[0].discard = 10

    state = endRound(state)

    equal(t, 0, int8(state.boards[0].score))
}

func Test_endGame_fullColour(t *testing.T) {
    state := newState()

    state.boards[0].grid[2][0] = true
    state.boards[0].grid[3][1] = true
    state.boards[0].grid[4][2] = true
    state.boards[0].grid[0][3] = true
    state.boards[0].grid[1][4] = true

    state = endGame(state)

    equal(t, 10, int8(state.boards[0].score))
}

func Test_endGame_fullColumn(t *testing.T) {
    state := newState()

    for i:=0; i<5; i++ {
        state.boards[1].grid[4][i] = true
    }

    state = endGame(state)

    equal(t, 7, int8(state.boards[1].score))
}

func Test_endGame_fullRow(t *testing.T) {
    state := newState()

    for i:=0; i<5; i++ {
        state.boards[1].grid[i][4] = true
    }

    state = endGame(state)

    equal(t, 2, int8(state.boards[1].score))
}

func Test_performMove_takeFromFirstPile(t *testing.T) {
    state := newState()

    state = appendPiles(state, 1)
    state.piles[0][0] = 2
    state.piles[0][1] = 2

    state = performMove(state, Move{
        pilesIndex: 0,
        colour: 1,
        rowIndex: 5,
    })

    equal(t, 2, state.piles[0][0])
    equal(t, 0, state.piles[0][1])
}

func Test_performMove_takeFromMiddlePile(t *testing.T) {
    state := newState()

    state = appendPiles(state, 3)
    state.piles[0][0] = 2
    state.piles[1][2] = 2
    state.piles[2][1] = 2

    state = performMove(state, Move{
        pilesIndex: 1,
        colour: 2,
        rowIndex: 5,
    })

    equal(t, 2, state.piles[0][0])
    equal(t, 2, state.piles[1][1])
}

func Test_performMove_putInDiscard(t *testing.T) {
    state := newState()

    state = appendPiles(state, 1)
    state.piles[0][0] = 2

    state = performMove(state, Move{
        pilesIndex: 0,
        colour: 0,
        rowIndex: 5,
    })

    equal(t, 2, state.boards[0].discard)
}

func Test_performMove_overfill(t *testing.T) {
    state := newState()

    state = appendPiles(state, 1)
    state.piles[0][2] = 3

    state = performMove(state, Move{
        pilesIndex: 0,
        colour: 2,
        rowIndex: 0,
    })

    equal(t, 2, state.boards[0].discard)
    equal(t, 2, state.boards[0].rows[0][0]) // colour
    equal(t, 1, state.boards[0].rows[0][1]) // index
}

func Test_performMove_underfill(t *testing.T) {
    state := newState()

    state = appendPiles(state, 1)
    state.piles[0][2] = 1

    state = performMove(state, Move{
        pilesIndex: 0,
        colour: 2,
        rowIndex: 2,
    })

    equal(t, 0, state.boards[0].discard)
    equal(t, 2, state.boards[0].rows[2][0]) // colour
    equal(t, 1, state.boards[0].rows[2][1]) // index
}

func appendPiles(state State, size int) State {
    for i:=0; i<size; i++ {
        var empty [5]int8
        state.piles = append(state.piles, empty)
    }
    return state
}