package xivapi

import (
	"fmt"
)

// Pagination holds the generic pagination info
// Only used on search
type Pagination struct {
	Page           int `json:"page"`
	PageNext       int `json:"page_next"`
	PagePrev       int `json:"page_prev"`
	PageTotal      int `json:"page_total"`
	Results        int `json:"results"`
	ResultsPerPage int `json:"results_per_page"`
	ResultsTotal   int `json:"results_total"`
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
