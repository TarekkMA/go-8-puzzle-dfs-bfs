package structs

type Stack struct {
	nodes []*Node
}

func (s *Stack) Push(node *Node) {
	s.nodes = append(s.nodes, node)
}

func (s *Stack) Pop() *Node {
	lastIndex := len(s.nodes) - 1
	n := s.nodes[lastIndex]
	s.nodes = s.nodes[:lastIndex]
	return n
}

func (s *Stack) IsEmpty() bool {
	return len(s.nodes) == 0
}
