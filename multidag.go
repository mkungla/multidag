package multidag

import (
	"errors"

	"github.com/google/uuid"
)

const (
	StatePending State = iota + 100
	StateSkipped
	StateInProgress
	StateSuccess
	StateFailure
)

const (
	PortIn PortType = iota + 200
	PortOut
)

type (
	State     uint8
	PortType  uint8
	Variables []Variable // do expose api .Has. Contains .Set .Update etc.

	Variable struct {
		Key   string `json:"key"` // key | name | property
		Value string `json:"value"`
		// Mask true idicates to mask variable value which presents can be known,
		// but the raw value should be masked in some cases. e.g secrets in logs.
		Mask    bool   `json:"mask,omitempty"`    // Should mask value e.g. ***
		Hint    string `json:"hint,omitempty"`    // hint used in config input
		Default string `json:"default,omitempty"` // default value when empty
	}

	Metadata struct {
		URI         URI    `json:"uri,omitempty"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}

	Authority struct {
		Namespace  string       `json:"namespace,omitempty"`
		Meta       Metadata     `json:"meta"`
		Workspaces []*Workspace `json:"workspaces,omitempty"`
	}

	// Workspace is collection of DAGs
	Workspace struct {
		Namespace string    `json:"namespace,omitempty"`
		Meta      Metadata  `json:"meta"`
		Vars      Variables `json:"vars,omitempty"` // variables available across all DAGs under this workspace
		DAGS      []*DAG    `json:"dags,omitempty"` // DAGs in this workspace
	}

	// DAG is directed acyclic graph of nodes.
	// Valid DAG consists of Node vertices and edges (arcs), with each edge directed from one vertex
	// to another, such that following those directions will never form a closed loop.
	// Nodes can have connections arrange such that satisfies above stament, but are free to have
	// any combination of dependencies within that constraint.
	// e.g.
	// 			[x,y,z]
	//      /  |  \
	// 	   /   |   \
	//  [x,y][x,z][y,z]
	// 	 | \/  |  \/ |
	// 	 | /\  |  /\ |
	// 	[x]  \ | /  [z]--[triggers]
	// 	  \   [y]   /
	// 	   \   |   /
	// 	 [trigger...]
	DAG struct {
		// locked bool     `json:"-"`
		Meta  Metadata `json:"meta"`
		State State    `json:"state,string"`
		Nodes []*Node  `json:"nodes,omitempty"`
	}

	Node struct {
		Meta  Metadata `json:"meta"`
		State State    `json:"state"`
		Ports []*Port  `json:"ports,omitempty"`
	}

	Port struct {
		Meta        Metadata `json:"meta"`
		Type        PortType `json:"type,string"`
		Connections []URI    `json:"connections,omitempty"`
	}

	// URI for MultiDag
	// [scheme]://[authority]/[workspace]/[node]/[port]?[query]#fragment
	URI struct {
		Scheme    string
		Authority string
		Workspace string
		DAG       uuid.UUID
		Node      uuid.UUID
		Port      uuid.UUID
		RawQuery  string // encoded query values, without '?'
		Fragment  string // fragment for references, without '#'
	}
)

var (
	ErrNamespaceEmpty = errors.New("namespace can not be empty")
)

func (s State) MarshalJSON() ([]byte, error) {
	var state string
	switch s {
	case StatePending:
		state = "\"pending\""
	case StateSkipped:
		state = "\"skipped\""
	case StateInProgress:
		state = "\"in-progress\""
	case StateSuccess:
		state = "\"success\""
	case StateFailure:
		state = "\"failure\""
	}
	return []byte(state), nil
}
