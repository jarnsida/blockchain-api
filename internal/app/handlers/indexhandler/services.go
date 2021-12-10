//go:generate mockgen -destination=./mocks.go -source=./services.go -package=indexhandler

package indexhandler

import (
	"context"
	"github.com/evt/blockchain-api/internal/e"
)

// IndexService is an index service.
type IndexService interface {
	GetIndex(ctx context.Context, id int64) (interface{}, e.Error)
}
