package model

import "github.com/matsunaga-project-2022/course-chat/internal/domain/valueobject"

type Workspace struct {
	ID   valueobject.WorkspaceID
	Name string
	Icon valueobject.Image
}
