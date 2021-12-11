//go:generate mockgen -destination=./mocks.go -source=./services.go -package=handlerindex

package handlerindex

import (
	"context"
	"github.com/evt/blockchain-api/internal/pkg/model"
)

// IndexService is an index service.
type IndexService interface {
	GetIndex(ctx context.Context, id int64) (*model.Index, error)
}
