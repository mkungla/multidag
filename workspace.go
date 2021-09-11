package multidag

import "github.com/google/uuid"

func (ws *Workspace) NewDAG(title string) *DAG {
	dag := &DAG{}
	uri := ws.Meta.URI
	uri.DAG, _ = uuid.NewRandom()
	dag.Meta.URI = uri
	dag.Meta.Title = title
	dag.State = StatePending
	ws.DAGS = append(ws.DAGS, dag)
	return dag
}
