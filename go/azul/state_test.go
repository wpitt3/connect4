package main

import (
	"testing"
	"strconv"
)


func Test_addTilesToPiles(t *testing.T) {
    state := newState()
     for i:=0; i<5; i++ {
        state.bag.tiles[i] = 4
    }

    // shuffle tiles into piles
    state = addTilesToPiles(state)

    var freq [5]int8
    for i:=0; i<6; i++ {
        for j:=0; j<5; j++ {
            freq[j] += state.piles[i][j]
        }
    }

    // 4 of each tiles in all piles
    // bag is empty
    for i:=0; i<5; i++ {
        equal(t, int8(0), state.bag.tiles[i])
        equal(t, int8(4), freq[i])
    }
}

func Test_addTilesToPiles_100tiles(t *testing.T) {
    state := newState()

    // shuffle tiles into piles
    state = addTilesToPiles(state)

    var freq [5]int8
    for i:=0; i<6; i++ {
             for j:=0; j<5; j++ {
                 freq[j] += state.piles[i][j]
             }
         }
    var sumFreq int8
    var sumInBag int8
    for i:=0; i<5; i++ {
        sumInBag += state.bag.tiles[i]
        sumFreq += freq[i]
    }
    equal(t, 80, sumInBag)
    equal(t, 20, sumFreq)

    for i:=0; i<5; i++ {
        freq[i] += state.bag.tiles[i]
        equal(t, int8(20), freq[i])
    }
}

func Test_find4_emptyBoard(t *testing.T) {
    var state State
    var piles [5]int8
    piles[0] = 2
    state.bag.tiles[0] = 1
    state.piles = append(state.piles, piles)
    state.boards[0].score = 3
    state.boards[1].discard = 4
    state.boards[0].rows[0][0] = 1
    state.boards[0].grid[0][0] = true

    newState := copyState(state)

    equal(t, int8(1), newState.bag.tiles[0])
    equal(t, int8(2), newState.piles[0][0])
    equal(t, int8(3), int8(newState.boards[0].score))
    equal(t, int8(4), newState.boards[1].discard)
    equal(t, int8(1), newState.boards[0].rows[0][0])
    equalBool(t, true, newState.boards[0].grid[0][0])

    newState.bag.tiles[0] = 2
    newState.piles[0][0] = 1
    newState.boards[0].score = 4
    newState.boards[1].discard = 5
    newState.boards[0].rows[0][0] = 2
    newState.boards[0].grid[0][0] = false

    equal(t, int8(1), state.bag.tiles[0])
    equal(t, int8(2), newState.bag.tiles[0])
    equal(t, int8(2), state.piles[0][0])
    equal(t, int8(1), newState.piles[0][0])
    equal(t, int8(3), int8(state.boards[0].score))
    equal(t, int8(4), int8(newState.boards[0].score))
    equal(t, int8(4), state.boards[1].discard)
    equal(t, int8(5), newState.boards[1].discard)
    equal(t, int8(1), state.boards[0].rows[0][0])
    equal(t, int8(2), newState.boards[0].rows[0][0])
    equalBool(t, true, state.boards[0].grid[0][0])
    equalBool(t, false, newState.boards[0].grid[0][0])
}


func equal(t *testing.T, x int8, y int8) {
    if x != y {
        t.Fatalf(strconv.Itoa(int(x)) + " should equal " + strconv.Itoa(int(y)))
    }
}

func equalBool(t *testing.T, x bool, y bool) {
    if x != y {
        t.Fatalf(" should equal " )
    }
}