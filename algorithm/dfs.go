package algorithm

import (
	"github.com/tarekkma/go-8-puzzle-dfs-bfs/structs"
)

func NewDFS(initial []int, final []int, limit int) *commonStruct {
	return &commonStruct{
		store:       &structs.Stack{},
		initialNode: structs.NewNodeWithMove(initial, nil, "Initial"),
		GoalNode:    structs.NewNode(final, nil),
		Limit:       limit,
		visitedList: make(map[int]bool),
	}
}
