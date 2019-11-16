package algorithm

import "github.com/tarekkma/go-8-puzzle-dfs-bfs/structs"

func CalcIntArrHashCode(arr []int) int {
	result := 1
	for i := 0; i < len(arr); i++ {
		result = 31*result + arr[i]
	}
	return result
}

type NodeStore interface {
	Add(node *structs.Node)
	Remove() *structs.Node
	IsEmpty() bool
	Clear()
}

type commonStruct struct {
	store       NodeStore
	initialNode *structs.Node
	visitedList map[int]bool
	LastNode    *structs.Node
	GoalNode    *structs.Node
	Limit       int
}

func (c *commonStruct) MarkAsVisited(n *structs.Node) {
	c.visitedList[n.CalcStateHashCode()] = true
}

func (c *commonStruct) IsGoalReached() bool {
	return c.LastNode.CalcStateHashCode() == c.GoalNode.CalcStateHashCode()
}

func (c *commonStruct) IsGoalReachedFor(n *structs.Node) bool {
	return n.CalcStateHashCode() == c.GoalNode.CalcStateHashCode()
}

func (c *commonStruct) IsVisited(n *structs.Node) bool {
	return c.visitedList[n.CalcStateHashCode()]
}

func (c *commonStruct) AddStatesFromLastNode(node *structs.Node) {
	node.DoAfterEachMovingState(func(node *structs.Node) {
		if c.IsVisited(node) == false {
			c.store.Add(node)
		}
	})
}

func (c *commonStruct) Excute() *structs.Node {
	levelMap := make(map[*structs.Node]int)

	c.store.Clear()
	c.LastNode = nil

	c.store.Add(c.initialNode)

	for !c.store.IsEmpty() {
		n := c.store.Remove()
		c.LastNode = n

		var level int
		if n == c.initialNode {
			level = 0
		} else {
			level = levelMap[n.Parent] + 1
		}
		levelMap[n] = level

		c.MarkAsVisited(c.LastNode)
		if c.IsGoalReachedFor(n) {
			break
		}

		if c.Limit == -1 || level < c.Limit {
			c.AddStatesFromLastNode(n)
		}
	}

	return c.LastNode
}
