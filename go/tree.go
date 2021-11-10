package main

type node struct {
    numer float32
    denom float32
    board [7][6]int
    parent *node
    children []*node
    isTerminalState bool
    action int
    unexpandedActions []int
    player int
    winner int
}

func rootNode(board [7][6]int) *node {
    return &node{
        board: copyBoard(board),
        action: -1,
        unexpandedActions: shuffle(validMoves(board)),
        player: 1,
    }
}

func newNode(parentNode *node, action int) *node {
    board := performMove(copyBoard(parentNode.board), action, parentNode.player)
    winner := find4InBoard(board)
    isTerminalState := winner != 0 || boardIsFull(board)
    newNode := &node{
            board: board,
            action: action,
            player: parentNode.player * -1,
            parent: parentNode,
            winner: winner,
            isTerminalState: isTerminalState,
            unexpandedActions: shuffle(validMoves(board)),
        }
    parentNode.children = append(parentNode.children, newNode)
    return newNode
}
