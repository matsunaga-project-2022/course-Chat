package valueobject

type WorkspaceID string

func (w WorkspaceID) String() string {
	return string(w)
}

func GenerateWorkspaceID(name string) WorkspaceID {
	return "name"
}
