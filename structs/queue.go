package structs

type Queue struct {
	nodes []*Node
}

func (q *Queue) Enqueue(node *Node) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) Dequeue() *Node {
	n := q.nodes[0]
	q.nodes = q.nodes[1:]
	return n
}

func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}
