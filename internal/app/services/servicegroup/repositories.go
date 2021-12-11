//go:generate mockgen -destination=./mocks.go -source=./repositories.go -package=servicegroup

package servicegroup

import (
	"context"
	"github.com/evt/blockchain-api/internal/pkg/models"
)

// Repository is a group repository.
type Repository interface {
	GetGroupIDs(context.Context) ([]int64, error)
	GetGroup(context.Context, int64) (*models.Group, error)
}
