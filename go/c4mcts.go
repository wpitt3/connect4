package main

import "C"

import (
	"reflect"
	"unsafe"
	"fmt"
)

// x := 0

//export Connect4MCTS
func Connect4MCTS(data *C.int, moves C.int) int {

	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(data)),
		Len:  int(42),
		Cap:  int(42),
	}
	slice := *(*[]int)(unsafe.Pointer(&header))

	var board [7][6]int
	for i:=0; i<7; i++ {
	    for j:=0; j<6; j++ {
	        board[i][j] = slice[i*6+j]
        }
	}

//     fmt.Printf("board %v\n", board)
//     fmt.Println(x)


    return findBestMove(board, int(moves))
}

func main() {}