package algorithm

import (
	"github.com/tarekkma/go-8-puzzle-dfs-bfs/structs"
)

func Bfs(initial []int, final []int) *structs.Node {
	queue := &structs.Queue{}
	initialNode := structs.NewNode(initial, nil)
	visitedList := make(map[int]bool)

	markAsVisited := func(n *structs.Node) {
		visitedList[n.CalcStateHashCode()] = true
	}
	isVisited := func(n *structs.Node) bool {
		return visitedList[n.CalcStateHashCode()]
	}
	addToQueueIfUnique := func(n *structs.Node) {
		if n == nil {
			return
		}
		if isVisited(n) {
			return
		}
		queue.Enqueue(n)
	}
	isGoalReached := func(n *structs.Node) bool {
		return CalcIntArrHashCode(final) == n.CalcStateHashCode()
	}

	queue.Enqueue(initialNode)

	var goalNode *structs.Node
	for !queue.IsEmpty() && goalNode == nil {
		currentNode := queue.Dequeue()
		markAsVisited(currentNode)
		if isGoalReached(currentNode) {
			goalNode = currentNode
			break
		}

		addToQueueIfUnique(currentNode.NodeMovedUp())
		addToQueueIfUnique(currentNode.NodeMovedDown())
		addToQueueIfUnique(currentNode.NodeMovedLeft())
		addToQueueIfUnique(currentNode.NodeMovedRight())
	}

	return goalNode
}
