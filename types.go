package xivapi

import (
	"encoding/json"
)

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

// JSONBoolNumber converts a number into a bool
type JSONBoolNumber bool

// UnmarshalJSON contains the conversion logic
func (jbn *JSONBoolNumber) UnmarshalJSON(bs []byte) error {
	var flag int
	if err := json.Unmarshal(bs, &flag); err != nil {
		return err
	}

	switch {
	case flag == 0:
		*jbn = false
	case flag != 1:
		return ErrUnexpectedType
	default:
		*jbn = true
	}

	return nil
}
