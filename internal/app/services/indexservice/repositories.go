//go:generate mockgen -destination=./mocks.go -source=./repositories.go -package=indexservice

package indexservice

import (
	"context"
)

// Repository is a repository.
type Repository interface {
	GetIndex(context.Context, int64) (interface{}, error)
}
