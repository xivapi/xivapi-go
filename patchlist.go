package xivapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// PatchEntries is a simple typed array
type PatchEntries []PatchEntry

// PatchEntry is a single PatchEntry of the PatchEntries list
type PatchEntry struct {
	TranslateableName

	Banner         string
	ExVersion      int
	ID             int
	IsExpansionRaw int   `json:"IsExpansion"`
	ReleaseDateRaw int64 `json:"ReleaseDate"`
	Version        json.Number
}

// ReleaseDate converts the raw unix timestamp of ReleaseDateRaw into a time.Time object
func (e PatchEntry) ReleaseDate() time.Time {
	return time.Unix(e.ReleaseDateRaw, 0)
}

// IsExpansion converts the IsExpansionRaw integer into a bool.
// This call treats -1 (no info available and not from this series) as false too
func (e PatchEntry) IsExpansion() bool {
	return e.IsExpansionRaw > 0
}

func (e PatchEntry) String() string {
	return fmt.Sprintf("Patch{ %v (%v) [%v] }", e.Name, e.Version, e.ReleaseDate())
}

// PatchList gets all known patches from the API
func (c *XIVAPI) PatchList() (PatchEntries, error) {
	uri, _ := url.Parse(fmt.Sprintf("%v%v", BaseURL, "/PatchList"))
	r := make(PatchEntries, 0, 25)

	if err := c.RequestJSON(MethodGet, uri, nil, &r); err != nil {
		return nil, err
	}

	return r, nil
}
