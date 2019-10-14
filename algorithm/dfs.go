package algorithm

import (
	"github.com/tarekkma/go-8-puzzle-dfs-bfs/structs"
)

func Dfs(initial []int, final []int) *structs.Node {
	stack := &structs.Stack{}
	initialNode := structs.NewNode(initial, nil)
	visitedList := make(map[int]bool)

	markAsVisited := func(n *structs.Node) {
		visitedList[n.CalcStateHashCode()] = true
	}
	isVisited := func(n *structs.Node) bool {
		return visitedList[n.CalcStateHashCode()]
	}
	addToStackIfUnique := func(n *structs.Node) {
		if n == nil {
			return
		}
		if isVisited(n) {
			return
		}
		stack.Push(n)
	}
	isGoalReached := func(n *structs.Node) bool {
		return CalcIntArrHashCode(final) == n.CalcStateHashCode()
	}

	stack.Push(initialNode)

	var goalNode *structs.Node
	for !stack.IsEmpty() && goalNode == nil {
		currentNode := stack.Pop()
		markAsVisited(currentNode)
		if isGoalReached(currentNode) {
			goalNode = currentNode
			break
		}

		addToStackIfUnique(currentNode.NodeMovedUp())
		addToStackIfUnique(currentNode.NodeMovedDown())
		addToStackIfUnique(currentNode.NodeMovedLeft())
		addToStackIfUnique(currentNode.NodeMovedRight())
	}
	return goalNode
}
