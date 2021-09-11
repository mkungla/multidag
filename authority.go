package multidag

func New(authority string) (*Authority, error) {
	ns := NewNamespace(authority)
	if ns == "" {
		return nil, ErrNamespaceEmpty
	}
	auth := &Authority{}
	auth.Namespace = ns
	auth.Meta.Title = ns
	auth.Meta.URI = URI{
		Scheme:    "multidag",
		Authority: ns,
	}
	return auth, nil
}

func (a *Authority) NewWorkspace(namespace string) (*Workspace, error) {
	ns := NewNamespace(namespace)
	if ns == "" {
		return nil, ErrNamespaceEmpty
	}

	uri := a.Meta.URI
	uri.Workspace = ns
	ws := &Workspace{
		Meta: Metadata{
			URI: uri,
		},
	}

	ws.Namespace = ns
	a.Workspaces = append(a.Workspaces, ws)
	return ws, nil
}
