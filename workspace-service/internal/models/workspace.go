package models

type Workspace struct {
	ID      WorkspaceID
	Name    string
	OwnerID string
	Icon    string
}

type WorkspaceID string

func (w WorkspaceID) String() string {
	return string(w)
}

func GenerateWorkspaceID(name string) WorkspaceID {
	return "name"
}
