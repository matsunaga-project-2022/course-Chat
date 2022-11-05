package repository

import "context"

type Manager interface {
	WithTransaction(ctx context.Context, fn func(tx Transaction) error) error
}

type Transaction interface {
	GetTx() any
}

type WorkspaceQueries interface {
	GetWorkspaceByID(id string) (WorkspaceQueries, error)
}
