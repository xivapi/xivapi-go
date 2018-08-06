package xivapi

// FileType is the type specifier for the type array
type FileType string

const (
	// FileData gets all additional data (e.g. members)
	FileData FileType = "data"

	// FileRecord gets the main record information
	FileRecord = "record"

	// FileMembers gets the members for the free company
	FileMembers = "members"
)
