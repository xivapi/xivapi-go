package xivapi

import (
	"fmt"
)

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
