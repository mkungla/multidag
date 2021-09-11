package multidag

import "github.com/google/uuid"

func (node *Node) NewPort(t PortType) *Port {
	port := &Port{}
	port.Type = t
	uri := node.Meta.URI
	uri.Port, _ = uuid.NewRandom()
	port.Meta.URI = uri
	node.Ports = append(node.Ports, port)
	return port
}
