package models

type Links struct {
	// Get the current page link
	Self string `json:"self"`
	// Get the next page link
	Next string `json:"next,omitempty"`
	// Get the last page link
	Last string `json:"last,omitempty"`
}
