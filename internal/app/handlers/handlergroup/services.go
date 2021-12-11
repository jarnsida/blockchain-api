//go:generate mockgen -destination=./mocks.go -source=./services.go -package=handlergroup

package handlergroup

import (
	"context"
	"github.com/evt/blockchain-api/internal/pkg/models"
)

// GroupService is a group service.
type GroupService interface {
	GetGroupIDs(ctx context.Context) ([]int64, error)
	GetGroup(ctx context.Context, id int64) (*models.Group, error)
}
