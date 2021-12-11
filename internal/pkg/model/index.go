package model

import "math/big"

// Index is a contract index.
type Index struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}
