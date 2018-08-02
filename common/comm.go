package common

// Error represent error type in each opearative error condition.
type Error struct {
	// Type:0: Client Error, Type: 1: System Error
	ErrorType        int
	ErrorDescription string
}
