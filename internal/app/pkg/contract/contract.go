package contract

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evt/blockchain-api/internal/contract"
	"math/big"
)

// Contract is a smart contract.
type Contract struct {
	*contract.Contract
}

// Bind binds to already deployed contract.
func Bind(contractAddress string, backend bind.ContractBackend) (*Contract, error) {
	if contractAddress == "" {
		return nil, errors.New("no contract address provided")
	}
	if backend == nil {
		return nil, errors.New("no contract backend provided")
	}

	// Get the contract address from hex string
	contractAddr := common.HexToAddress(contractAddress)

	ctr, err := contract.NewContract(contractAddr, backend)
	if err != nil {
		return nil, fmt.Errorf("contract.NewContract failed: %w", err)
	}

	return &Contract{ctr}, nil
}

func (ec *Contract) GetGroupIDs(ctx context.Context) ([]int64, error) {
	groupIDs, err := ec.Contract.GetGroupIds(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("ctr.GetGroupIds failed: %w", err)
	}

	result := make([]int64, 0, len(groupIDs))
	for _, groupID := range groupIDs {
		result = append(result, groupID.Int64())
	}

	return result, nil
}

func (ec *Contract) GetGroup(ctx context.Context, id int64) (interface{}, error) {
	group, err := ec.Contract.GetGroup(&bind.CallOpts{Context: ctx}, big.NewInt(id))
	if err != nil {
		return nil, fmt.Errorf("ctr.GetGroup failed: %w", err)
	}

	return group, nil
}
