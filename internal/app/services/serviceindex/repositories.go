//go:generate mockgen -destination=./mocks.go -source=./repositories.go -package=serviceindex

package serviceindex

import (
	"context"
	"github.com/evt/blockchain-api/internal/pkg/models"
)

// Repository is an index repository.
type Repository interface {
	GetIndex(context.Context, int64) (*models.Index, error)
}
