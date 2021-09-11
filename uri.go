package multidag

import (
	"net/url"

	"github.com/google/uuid"
)

func (u *URI) String() string {
	host := u.Authority
	var path string
	if len(u.Workspace) > 0 {
		path = u.Workspace
		if u.DAG != uuid.Nil {
			path += "/" + u.DAG.String()
			if u.Node != uuid.Nil {
				path += "/" + u.Node.String()
				if u.Port != uuid.Nil {
					path += "/" + u.Port.String()
				}
			}
		}
	}

	url := url.URL{
		Scheme: "multidag",
		Host:   host,
		Path:   path,
	}
	return url.String()
}
