package model

import "math/big"

// Group is a contract group
type Group struct {
	Name    string
	Indexes []*big.Int
}
