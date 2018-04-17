package gonet

// Error Error message object with error code and arguments for client-side message resolving
type Error struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Args    []interface{} `json:"args"`
}
