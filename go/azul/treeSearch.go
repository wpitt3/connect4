package main

import (
	"math"
// 	"fmt"
    "math/rand"
    "time"
)


// func findBestMove(state State, numberOfMoves int) int {
//     rootNode := rootNode(state)
//     for i:=0; i<numberOfMoves; i++ {
//         leafNode := selectLeafNode(rootNode)
//         result := simulateGame(leafNode.board, leafNode.player)
//         var score float32
//         if result == leafNode.player {
//             score = 0.0
//         } else if result * -1 == leafNode.player {
//             score = 1.0
//         } else {
//             score = 0.5
//         }
//         currentNode := leafNode
//         for currentNode != rootNode {
//             currentNode.numer += score
//             currentNode.denom += 1
//             score = 1.0 - score
//             currentNode = currentNode.parent
//         }
//         currentNode.numer += score
//         currentNode.denom += 1
//     }
//
//     sortChildren(rootNode.children)
//
//     bestMove := -1
//     bestScore := float32(0.0)
//     for i:=0; i<len(rootNode.children); i++ {
//         child := rootNode.children[i]
//         fmt.Println(child.numer/child.denom)
//         if (child.numer/child.denom) > bestScore {
//             bestScore = (child.numer/child.denom)
//             bestMove = child.action
//         }
//     }
//     fmt.Println()
//     return bestMove
// }

func simulateGame(oldState State, semiTerminalState bool, isTerminalState bool) int {
    rand.Seed(time.Now().UTC().UnixNano())
    state := copyState(oldState)
    if (semiTerminalState) {
        state = addTilesToPiles(state)
    }

    done := !isTerminalState

    for(!done) {
        moves := validMoves(state)
        move := moves[rand.Intn(len(moves))]
        state = performMove(state, move)

        // if end of round, refill piles
        // if end of game, done is true

        endOfRound := isRoundOver(state)

        if (endOfRound) {
            state = endRound(state)
            if (!isGameOver(state)) {
                state = endGame(state)
                done = true
            } else {
                state = addTilesToPiles(state)
            }
        }
    }
    return result
}

// with semiTerminalStates they should still be selectable but not expandable
func selectLeafNode(rootNode *node) *node {
    currentNode := rootNode
    c := float32(1.414)
    for (!currentNode.isTerminalState && !currentNode.isSemiTerminalState) {

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

type node struct {
    numer float32
    denom float32
    state State
    parent *node
    children []*node
    isSemiTerminalState bool
    isTerminalState bool
    action Move
    unexpandedActions []Move
}

func rootNode(state State) *node {
    return &node{
        state: copyState(state),
        unexpandedActions: shuffleMoves(findValidMoves(state)),
    }
}

func newNode(parentNode *node, move Move) *node {
    state := performMove(copyState(parentNode.state), move)
    isSemiTerminalState := isRoundOver(state)
    if (isSemiTerminalState) {
        state = endRound(state)
    }
    isTerminalState := isSemiTerminalState && isGameOver(state)
    if (isTerminalState) {
        state = endGame(state)
    }

    newNode := &node{
        state: state,
        action: move,
        parent: parentNode,
        isSemiTerminalState: isSemiTerminalState,
        isTerminalState: isTerminalState,
        unexpandedActions: shuffleMoves(findValidMoves(state)),
    }
    parentNode.children = append(parentNode.children, newNode)
    return newNode
}

