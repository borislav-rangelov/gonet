package response

import "github.com/borislav-rangelov/gonet"

// Error Error message wrapper
type Error struct {
	Code   int                       `json:"code"`
	Errors *map[string][]gonet.Error `json:"errors"`
}
