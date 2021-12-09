package groupservice

import (
	"context"
	"github.com/evt/blockchain-api/internal/e"
)

// GroupService is a group service.
type GroupService struct {
	contract SmartContract
}

// New creates a new group service.
func New(contract SmartContract) *GroupService {
	return &GroupService{
		contract: contract,
	}
}

// GetGroups returns smart contract groups.
func (svc *GroupService) GetGroupIDs(ctx context.Context) ([]int64, e.Error) {
	groupIDs, err := svc.contract.GetGroupIDs(ctx)
	if err != nil {
		return nil, e.NewInternalf("failed getting smart contract groups: %s", err)
	}

	return groupIDs, nil
}

// GetGroups returns smart contract groups.
func (svc *GroupService) GetGroup(ctx context.Context, id int64) (interface{}, e.Error) {
	group, err := svc.contract.GetGroup(ctx, id)
	if err != nil {
		return nil, e.NewInternalf("failed getting smart contract groups: %s", err)
	}

	return group, nil
}
