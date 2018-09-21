package xivapi

// Version is the package version
const Version = "0.1.0"

//State Shows current status of a object on XIVAPI
type State int

//Collection of all states of a object on request
const (
	StateNone State = iota
	StateAdding
	StateCached
	StateNotFound
	StateBlacklist
	StatePrivate
)
