package multidag

import "github.com/google/uuid"

func (dag *DAG) NewNode(title string) *Node {
	node := &Node{}
	uri := dag.Meta.URI
	uri.Node, _ = uuid.NewRandom()
	node.Meta.URI = uri
	node.Meta.Title = title
	node.State = StatePending
	dag.Nodes = append(dag.Nodes, node)
	return node
}
