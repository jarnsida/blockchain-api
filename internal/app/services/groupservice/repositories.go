//go:generate mockgen -destination=./mocks.go -source=./repositories.go -package=groupservice

package groupservice

import (
	"context"
)

// SmartContract is a smart contract.
type SmartContract interface {
	GetGroupIDs(context.Context) ([]int64, error)
	GetGroup(context.Context, int64) (interface{}, error)
}
