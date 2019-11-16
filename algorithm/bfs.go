package algorithm

import (
	"github.com/tarekkma/go-8-puzzle-dfs-bfs/structs"
)

type bfs struct {
	commonStruct
}

func NewBFS(initial []int, final []int) *commonStruct {
	return &commonStruct{
		store:       &structs.Queue{},
		initialNode: structs.NewNodeWithMove(initial, nil, "Initial"),
		GoalNode:    structs.NewNode(final, nil),
		Limit:       -1,
		visitedList: make(map[int]bool),
	}
}
