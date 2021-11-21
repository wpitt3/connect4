package main

import (
    "math/rand"
    "time"
)

func performMove(state State, move Move) State {
    pileIndex := move.pilesIndex
    var colourCount int8
    if (pileIndex > 0) {
        pile := state.piles[pileIndex]
        colourCount = pile[move.colour]
        for colour:=int8(0); colour<5; colour++ {
            if (colour != move.colour) {
                state.piles[0][colour] += pile[colour]
            }
        }
        state.piles = append(state.piles[:pileIndex], state.piles[(pileIndex+int8(1)):]...)
    } else {
        colourCount = state.piles[0][move.colour]
        state.piles[0][move.colour] = 0
    }

    p := state.nextPlayer
    if (move.rowIndex == 5) {
        state.boards[p].discard += colourCount
    } else {
        board := state.boards[p]
        row := move.rowIndex
        if (colourCount + board.rows[row][1] > row + 1) {
            remainder := (colourCount + board.rows[row][1]) - (row + 1)
            state.boards[p].rows[row][1] = row + 1
            state.boards[p].rows[row][0] = move.colour
            state.boards[p].discard += remainder
        } else {
            state.boards[p].rows[row][1] += colourCount
            state.boards[p].rows[row][0] = move.colour
        }
    }

    return state
}

func findValidMoves(state State) []Move {
    var moves []Move

    validColumnColours := findValidColumnColours(state)

    for pileIndex:=int8(0); pileIndex<int8(len(state.piles)); pileIndex++ {
        for colour:=int8(0); colour<5; colour++ {
            if (state.piles[pileIndex][colour] > 0) {
                moves = append(moves, Move{
                    pilesIndex: pileIndex,
                    colour: colour,
                    rowIndex: 5,
                })
                for rowIndex:=int8(0); rowIndex<5; rowIndex++ {
                    if (validColumnColours[rowIndex*5+colour]) {
                        moves = append(moves, Move{
                            pilesIndex: pileIndex,
                            colour: colour,
                            rowIndex: rowIndex,
                        })
                    }
                }
            }
        }
    }

    return moves
}

// returns a set of valid moves with row * 5 + colour
func findValidColumnColours(state State) map[int8]bool {
    validColumnColours := make(map[int8]bool)
    board := state.boards[state.nextPlayer]
    for i:=int8(0); i<5; i++ {
        if (board.rows[i][1] > 0) {
            if (board.rows[i][1] <= i) {
                validColumnColours[i * 5 + board.rows[i][0]] = true
            }
        } else {
            for j:=int8(0); j<5; j++ {
                if (!board.grid[j][i]) {
                    validColumnColours[i * 5 + findGridColour(j, i)] = true
                }
            }
        }
    }
    return validColumnColours
}

func isRoundOver(state State) bool {
    for i:=0; i<len(state.piles); i++ {
        for j:=0; j<5; j++ {
            if (state.piles[i][j] != 0) {
                return false
            }
        }
    }
    return true
}

func isGameOver(state State) bool {
    for boardIndex:=0; boardIndex<2; boardIndex++ {
        for row:=0; row<5; row++ {
            allFull := true
            for j:=0; j<5; j++ {
                if (!state.boards[boardIndex].grid[j][row]) {
                    allFull = false
                }
            }
            if (allFull) {
                return true
            }
        }
    }
    return false
}

func endGame(state State) State {
    for boardIndex:=0; boardIndex<2; boardIndex++ {
        board := state.boards[boardIndex]
        for colour:=int8(0); colour<5; colour++ {
            full := true
            for row:=int8(0); row<5; row++ {
                if (!board.grid[findGridColumn(colour, row)][row]) {
                    full = false
                }
            }
            if (full) {
                state.boards[boardIndex].score += 10
            }
        }
        for row:=int8(0); row<5; row++ {
            full := true
            for column:=int8(0); column<5; column++ {
                if (!board.grid[column][row]) {
                    full = false
                }
            }
            if (full) {
                state.boards[boardIndex].score += 2
            }
        }
        for column:=int8(0); column<5; column++ {
            full := true
            for row:=int8(0); row<5; row++ {
                if (!board.grid[column][row]) {
                    full = false
                }
            }
            if (full) {
                state.boards[boardIndex].score += 7
            }
        }
    }
    return state
}

func endRound(state State) State {
    for boardIndex:=0; boardIndex<2; boardIndex++ {
        board := state.boards[boardIndex]
        for row:=int8(0); row<5; row++ {
            if (board.rows[row][1] == row + 1) {
                column := findGridColumn(board.rows[row][0], row)
                state.boards[boardIndex].rows[row][1] = 0
                state.boards[boardIndex].grid[column][row] = true
                state.boards[boardIndex].score += countAdjacent(board, column, row)
            }
        }
        if (board.discard < 3) {
            state.boards[boardIndex].score -= int(board.discard)
        } else if (board.discard < 6) {
            state.boards[boardIndex].score -= int((board.discard - 1) * 2)
        } else {
            discard := min(int(board.discard), 7)
            state.boards[boardIndex].score -= int((discard-5)*3+8)
        }
        state.boards[boardIndex].score = max(state.boards[boardIndex].score, 0)

    }
    return state
}

func countAdjacent(board Board, column int8, row int8) int {
    height := 1
    width := 1
    x := row + int8(1)
    for(x<5 && board.grid[column][x]) {
        height += 1
        x++
    }
    x = row - int8(1)
    for(x>=0 && board.grid[column][x]) {
        height += 1
        x--
    }

    y := column + int8(1)
    for(y<5 && board.grid[y][row]) {
        y++
        width += 1
    }
    y = column - int8(1)
    for(y>=0 && board.grid[y][row]) {
        y--
        width += 1
    }
    if (height == 1 || width == 1) {
        return height * width
    }
    return height + width
}

func findGridColour(x int8, row int8) int8 {
    return (5 + x - row) % 5
}

func findGridColumn(colour int8, row int8) int8 {
    return (colour + row) % 5
}

func min(a int, b int) int {
    if (a > b) {
        return b
    }
    return a
}

func max(a int, b int) int {
    if (a < b) {
        return b
    }
    return a
}

func shuffleMoves(list []Move) []Move {
    rand.Seed(time.Now().UTC().UnixNano())
    for i := len(list); i>0; i-- {
        ri := rand.Intn(i)
        temp := list[ri]
        list[ri] = list[i-1]
        list[i-1] = temp
    }
    return list
}