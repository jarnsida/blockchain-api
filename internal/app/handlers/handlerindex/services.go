//go:generate mockgen -destination=./mocks.go -source=./services.go -package=indexhandler

package handlerindex

import (
	"context"
)

// IndexService is an index service.
type IndexService interface {
	GetIndex(ctx context.Context, id int64) (interface{}, error)
}
