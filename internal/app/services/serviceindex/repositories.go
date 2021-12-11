//go:generate mockgen -destination=./mocks.go -source=./repositories.go -package=serviceindex

package serviceindex

import (
	"context"
)

// Repository is an index repository.
type Repository interface {
	GetIndex(context.Context, int64) (interface{}, error)
}
