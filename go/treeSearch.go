package main

import (
	"math"
// 	"fmt"
)

func findBestMove(board [7][6]int, moves int) int {
    rootNode := rootNode(board)
    for i:=0; i<moves; i++ {
        leafNode := selectLeafNode(rootNode)
        result := simulateGame(leafNode.board, leafNode.player)
        var score float32
        if result == leafNode.player {
            score = 0.0
        } else if result * -1 == leafNode.player {
            score = 1.0
        } else {
            score = 0.5
        }
        currentNode := leafNode
        for currentNode != rootNode {
            currentNode.numer += score
            currentNode.denom += 1
            score = 1.0 - score
            currentNode = currentNode.parent
        }
        currentNode.numer += score
        currentNode.denom += 1
    }

    sortChildren(rootNode.children)

    bestMove := -1
    bestScore := float32(0.0)
    for i:=0; i<len(rootNode.children); i++ {
        child := rootNode.children[i]
//         fmt.Println(child.numer/child.denom)
        if (child.numer/child.denom) > bestScore {
            bestScore = (child.numer/child.denom)
            bestMove = child.action
        }
    }
//     fmt.Println()
    return bestMove
}

func selectLeafNode(rootNode *node) *node {
    currentNode := rootNode
    c := float32(1.414)
    for (!currentNode.isTerminalState) {

        if len(currentNode.unexpandedActions) > 0 {
            newNode := newNode(currentNode, currentNode.unexpandedActions[0])
            currentNode.unexpandedActions = currentNode.unexpandedActions[1:]
            return newNode
        }
        currentNode = findBestChild(currentNode, c)
    }

    return currentNode
}

func findBestChild(parent *node, c float32) *node {
    logTotalParent := math.Log(float64(parent.denom))
    var maxScore float32 = 0.0
    var bestChild *node = parent.children[0]
    for i:=0; i<len(parent.children); i++ {
        child := parent.children[i]
        score := exploreFunction(child.numer, child.denom, logTotalParent, c)
        if score > maxScore {
            maxScore = score
            bestChild = child
        }
    }
    return bestChild
}

func exploreFunction(wins float32, total float32, logTotalParent float64, c float32) float32 {
    return wins/total + c * float32(math.Sqrt(logTotalParent/float64(total)))
}

func performMove(board [7][6]int, column int, player int) [7][6]int {
    for i:=0; i<6; i++ {
        if board[column][i] == 0 {
            board[column][i] = player
            return board
        }
    }
    return board
}

func validMove(board [7][6]int, column int) bool {
    return board[column][5] == 0
}

func validMoves(board [7][6]int) []int {
    validMoves := make([]int, 0)
    for i:=0; i<7; i++ {
        if validMove(board, i) {
            validMoves = append(validMoves, i)
        }
    }
    return validMoves
}

func boardIsFull(board [7][6]int) bool {
    for i:=0; i<7; i++ {
        if board[i][5] == 0 {
            return false
        }
    }
    return true
}

func find4InBoard(board [7][6]int) int {
    for i:=0; i<7; i++ {
        for j:=0; j<3; j++ {
            var sum = 0
            for k:=0; k<4; k++ {
                sum += board[i][j+k]
            }
            if abs(sum) == 4 {
                return sum/4
            }
        }
    }
    for i:=0; i<4; i++ {
        for j:=0; j<6; j++ {
            var sum = 0
            for k:=0; k<4; k++ {
                sum += board[i+k][j]
            }
            if abs(sum) == 4 {
                return sum/4
            }
        }
    }
    for i:=0; i<4; i++ {
            for j:=0; j<3; j++ {
                var sumNE = 0
                var sumNW = 0
                for k:=0; k<4; k++ {
                    sumNE += board[i+k][j+k]
                    sumNW += board[i+k][3+j-k]
                }
                if abs(sumNE) == 4 {
                    return sumNE/4
                }
                if abs(sumNW) == 4 {
                    return sumNW/4
                }
            }
        }

    return 0
}