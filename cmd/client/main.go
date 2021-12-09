package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evt/blockchain-api/contract"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Connect to a node
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/101df4c66df249d1ae14489c8af62eb3")
	if err != nil {
		return fmt.Errorf("ethclient.Dial failed: %w", err)
	}

	// Get the contract address from hex string
	contractAddr := common.HexToAddress("0x4f7f1380239450AAD5af611DB3c3c1bb51049c29")

	// Bind to an already deployed contract
	ctr, err := contract.NewContract(contractAddr, client)
	if err != nil {
		return fmt.Errorf("contract.NewContract failed: %w", err)
	}

	ctx := context.Background()
	callOpts := &bind.CallOpts{Context: ctx}

	groupIDs, err := ctr.GetGroupIds(callOpts)
	if err != nil {
		return fmt.Errorf("ctr.GetGroupIds failed: %w", err)
	}

	for _, groupID := range groupIDs {
		group, err := ctr.GetGroup(callOpts, groupID)
		if err != nil {
			return fmt.Errorf("ctr.GetGroupIds failed: %w", err)
		}

		spew.Dump(group)

		for _, indexID := range group.Indexes {
			index, err := ctr.GetIndex(callOpts, indexID)
			if err != nil {
				return fmt.Errorf("ctr.GetGroupIds failed: %w", err)
			}

			spew.Dump(index)
		}
	}

	return nil
}
