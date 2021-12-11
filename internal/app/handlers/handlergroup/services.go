//go:generate mockgen -destination=./mocks.go -source=./services.go -package=grouphandler

package handlergroup

import (
	"context"
)

// GroupService is a group service.
type GroupService interface {
	GetGroupIDs(ctx context.Context) ([]int64, error)
	GetGroup(ctx context.Context, id int64) (interface{}, error)
}
