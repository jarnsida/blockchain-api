package models

// Error is a generic error returned in API response.
// swagger:model error
type Error struct {
	// The error message
	// required: true
	Error string `json:"error"`
}
