package xivapi

import (
	"fmt"
)

// Pagination holds the generic pagination info
// Only used on search
type Pagination struct {
	Page           int
	PageNext       int
	PagePrev       int
	PageTotal      int
	Results        int
	ResultsPerPage int
	ResultsTotal   int
}

func (p Pagination) String() string {
	return fmt.Sprintf("Pagination{%d / %d (%d)}", p.Page, p.PageTotal, p.ResultsTotal)
}

// TranslateableName holds all different names for an object
// Only used on search so far, but used in different structs
type TranslateableName struct {
	Name   string `json:"Name"`
	NameCN string `json:"Name_cn"`
	NameDE string `json:"Name_de"`
	NameEN string `json:"Name_en"`
	NameFR string `json:"Name_fr"`
	NameJA string `json:"Name_ja"`
	NameKR string `json:"Name_kr"`
}
