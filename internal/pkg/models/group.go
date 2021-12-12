package models

// Group is a contract group.
// swagger:model group
type Group struct {
	// example: DeFi Indexes
	Name string
	// example: [0,1,2,3,4,5]
	Indexes []int64
}

// Groups holds a list of group IDs.
// swagger:model groupIDs
type Groups struct {
	// example: [12,13]
	GroupIDs []int64 `json:"group_ids"`
}
