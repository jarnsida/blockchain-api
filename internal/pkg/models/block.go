package models

import "github.com/ethereum/go-ethereum/core/types"

// Block is an Ethereum block as returned by Ethereum client.
type Block struct {
	Header *types.Header `json:",omitempty"`
	Body   *types.Body   `json:",omitempty"`
}
