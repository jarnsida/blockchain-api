package models

// Index is a contract index.
// swagger:model index
type Index struct {
	// example: DeFi Index (1)
	Name string
	// example: 150000000000000000
	EthPriceInWei int64
	// example: 9500
	UsdPriceInCents int64
	// example: 250000000
	UsdCapitalization int64
	// example: -45
	PercentageChange int64
}
