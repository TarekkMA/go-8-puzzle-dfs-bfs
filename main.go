package main

import (
	"fmt"
	"time"

	"github.com/tarekkma/go-8-puzzle-dfs-bfs/algorithm"
	"github.com/tarekkma/go-8-puzzle-dfs-bfs/structs"
)

func main() {
	initial := []int{
		2, 8, 3,
		1, 6, 4,
		7, 0, 5,
	}
	final := []int{
		1, 2, 3,
		8, 0, 4,
		7, 6, 5,
	}
	var bfsRes *structs.Node
	var dfsRes *structs.Node

	measureFunctionTime("BFS", func() {
		bfsRes = algorithm.NewBFS(initial, final).Excute()
	})
	measureFunctionTime("DFS LIMITED DEAPTH", func() {
		dfsRes = algorithm.NewDFS(initial, final, -1).Excute()
	})
	if dfsRes.CalcStateHashCode() == bfsRes.CalcStateHashCode() {
		fmt.Println("Same result")
	} else {
		fmt.Println("Different results.")
	}
	fmt.Println()
	PrintMoves("DFS", dfsRes, false, false)
	PrintMoves("BFS", bfsRes, false, false)
}

func measureFunctionTime(name string, f func()) {
	fmt.Printf("%s is starting\n", name)
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("%s finished in %s\n", name, elapsed)
}

func PrintMoves(name string, n *structs.Node, shouldPrintMovesList bool, shouldPrintVisualTrace bool) {
	moves := n.GetMoves()
	fmt.Printf("%s took %d to be solved\n", name, len(moves))
	if shouldPrintMovesList {
		fmt.Println(moves)
	}
	if shouldPrintVisualTrace {
		n.PrintTrace()
	}
}
