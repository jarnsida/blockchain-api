//go:generate mockgen -destination=./mocks.go -source=./repositories.go -package=groupservice

package groupservice

import (
	"context"
)

// Repository is a group repository.
type Repository interface {
	GetGroupIDs(context.Context) ([]int64, error)
	GetGroup(context.Context, int64) (interface{}, error)
}
