package models

import "github.com/ethereum/go-ethereum/core/types"

// Block is an Ethereum block as returned by Ethereum client.
// swagger:model block
type Block struct {
	Header *types.Header `json:",omitempty"`
	Body   *types.Body   `json:",omitempty"`
}

// BlockNoBody is an Ethereum block with header only as returned by Ethereum client.
// swagger:model blockNoBody
type BlockNoBody struct {
	Header *types.Header `json:",omitempty"`
}
