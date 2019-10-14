package structs

import "fmt"

type Node struct {
	Parent   *Node
	Children []*Node
	Move     string

	GameBoard []int
}

func NewNode(gameBoard []int, parent *Node) *Node {
	return &Node{
		Parent:    parent,
		GameBoard: gameBoard,
		Children:  make([]*Node, 0),
		Move:      "initial",
	}
}

func NewNodeWithMove(gameBoard []int, parent *Node, move string) *Node {
	return &Node{
		Parent:    parent,
		GameBoard: gameBoard,
		Children:  make([]*Node, 0),
		Move:      move,
	}
}

func (n *Node) CalcStateHashCode() int {
	result := 1
	for i := 0; i < len(n.GameBoard); i++ {
		result = 31*result + n.GameBoard[i]
	}
	return result
}

func (n *Node) GetMoves() []string {
	moves := []string{}
	var recursive func(node *Node)
	recursive = func(node *Node) {
		if node.Parent != nil {
			recursive(node.Parent)
		}
		moves = append(moves, node.Move)
	}
	recursive(n)
	return moves
}

func (n *Node) PrintTrace() {
	if n.Parent != nil {
		n.Parent.PrintTrace()
	}
	fmt.Printf("%s\n", n.Move)
	n.Print()
	fmt.Println()
}

func (n *Node) Print() {
	fmt.Printf(`%v %v %v
%v %v %v
%v %v %v
`,
		n.GameBoard[0],
		n.GameBoard[1],
		n.GameBoard[2],
		n.GameBoard[3],
		n.GameBoard[4],
		n.GameBoard[5],
		n.GameBoard[6],
		n.GameBoard[7],
		n.GameBoard[8],
	)
}

func (n *Node) findZeroIndex() int {
	zIndex := -1
	for i := 0; i < len(n.GameBoard); i++ {
		if n.GameBoard[i] == 0 {
			zIndex = i
		}
	}
	if zIndex == -1 {
		panic("Zero not found in the game board")
	}
	return zIndex
}

func (n *Node) NodeMovedUp() *Node {
	zPos := n.findZeroIndex()

	tmpGB := make([]int, len(n.GameBoard))
	copy(tmpGB, n.GameBoard)

	if zPos != 0 && zPos != 1 && zPos != 2 {
		tmpGB[zPos], tmpGB[zPos-3] = tmpGB[zPos-3], tmpGB[zPos]
		return NewNodeWithMove(tmpGB, n, "Up")
	}
	return nil
}

func (n *Node) NodeMovedDown() *Node {
	zPos := n.findZeroIndex()

	tmpGB := make([]int, len(n.GameBoard))
	copy(tmpGB, n.GameBoard)

	if zPos != 6 && zPos != 7 && zPos != 8 {
		tmpGB[zPos], tmpGB[zPos+3] = tmpGB[zPos+3], tmpGB[zPos]
		return NewNodeWithMove(tmpGB, n, "Down")
	}
	return nil
}

func (n *Node) NodeMovedRight() *Node {
	zPos := n.findZeroIndex()

	tmpGB := make([]int, len(n.GameBoard))
	copy(tmpGB, n.GameBoard)

	if zPos != 2 && zPos != 5 && zPos != 8 {
		tmpGB[zPos], tmpGB[zPos+1] = tmpGB[zPos+1], tmpGB[zPos]
		return NewNodeWithMove(tmpGB, n, "Right")
	}
	return nil
}

func (n *Node) NodeMovedLeft() *Node {
	zPos := n.findZeroIndex()

	tmpGB := make([]int, len(n.GameBoard))
	copy(tmpGB, n.GameBoard)

	if zPos != 0 && zPos != 3 && zPos != 6 {
		tmpGB[zPos], tmpGB[zPos-1] = tmpGB[zPos-1], tmpGB[zPos]
		return NewNodeWithMove(tmpGB, n, "Left")
	}
	return nil
}
