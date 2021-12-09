//go:generate mockgen -destination=./mocks.go -source=./services.go -package=grouphandler

package grouphandler

import (
	"context"
	"github.com/evt/blockchain-api/internal/e"
)

// GroupService is a group service.
type GroupService interface {
	GetGroupIDs(ctx context.Context) ([]int64, e.Error)
	GetGroup(ctx context.Context, id int64) (interface{}, e.Error)
}
