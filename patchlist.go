package xivapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type PatchEntries []PatchEntry

type PatchEntry struct {
	TranslateableName

	Banner         string
	ExVersion      int
	ID             int
	IsExpansionRaw int   `json:"IsExpansion"`
	ReleaseDateRaw int64 `json:"ReleaseDate"`
	Version        json.Number
}

func (e PatchEntry) ReleaseDate() time.Time {
	return time.Unix(e.ReleaseDateRaw, 0)
}

func (e PatchEntry) IsExpansion() bool {
	return e.IsExpansionRaw > 0
}

func (e PatchEntry) String() string {
	return fmt.Sprintf("Patch{ %v (%v) [%v] }", e.Name, e.Version, e.ReleaseDate())
}

func (c *XIVAPI) PatchList() (PatchEntries, error) {
	uri, _ := url.Parse(fmt.Sprintf("%v%v", BaseURL, "/PatchList"))
	r := make(PatchEntries, 0, 25)

	if err := c.RequestJSON(MethodGet, uri, nil, &r); err != nil {
		return nil, err
	}

	return r, nil
}
