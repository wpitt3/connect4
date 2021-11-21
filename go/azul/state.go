package main

import (
    "math/rand"
    "time"
)

type Move struct {
    pilesIndex int8
    colour int8
    rowIndex int8
}

func newState() State {
    state := State{firstPlayerNextRound: -1}
    for i:=0; i<5; i++ {
        state.bag.tiles[i] = 20
    }
    return state
}

type State struct {
    // 2 player only at the moment
    // piles[0] is centre of rejected tiles
    piles [][5]int8
    boards [2]Board
    bag Bag
    nextPlayer int8
    firstPlayerNextRound int8
}

type Board struct {
    score int
    discard int8
    // rows has colour and then size
    rows [5][2]int8
    grid [5][5]bool
}

type Bag struct {
    tiles [5]int8
}

func copyState(state State) State {
    var newState = State{}
    for i:=0; i<5; i++ {
        newState.bag.tiles[i] = state.bag.tiles[i]
    }

    for i:=0; i<len(state.piles); i++ {
        var pile [5]int8
        for j:=0; j<5; j++ {
            pile[j] = state.piles[i][j]
        }
        newState.piles = append(newState.piles, pile)
    }

    for i:=0; i<2; i++ {
        newState.boards[i].score = state.boards[i].score
        newState.boards[i].discard = state.boards[i].discard
        for j:=0; j<5; j++ {
            for k:=0; k<5; k++ {
                newState.boards[i].grid[j][k] = state.boards[i].grid[j][k]
            }
            for k:=0; k<2; k++ {
                newState.boards[i].rows[j][k] = state.boards[i].rows[j][k]
            }
        }
    }
    return newState
}

func shuffleBag(list []int8) []int8 {
    rand.Seed(time.Now().UTC().UnixNano())
    for i := len(list); i>0; i-- {
        ri := rand.Intn(i)
        temp := list[ri]
        list[ri] = list[i-1]
        list[i-1] = temp
    }
    return list
}

func addTilesToPiles(state State) State {
    var tiles []int8
    if len(state.bag.tiles) < 5*4 {
        // refill bag
        // count tiles on boards rows + grid
    }
    for i:=int8(0); i<5; i++ {
        for j:=0; j<int(state.bag.tiles[i]); j++ {
            tiles = append(tiles, i)
        }
    }
    shuffleBag(tiles)
    var empty [5]int8
    state.piles = append(state.piles, empty)
    for i:=0; i<5; i++ {
        var pile [5]int8
        for j:=0; j<4; j++ {
            tile := tiles[i*4+j]
            state.bag.tiles[tile] -= 1
            pile[tile] += 1
        }
        state.piles = append(state.piles, pile)
    }

    return state
}

