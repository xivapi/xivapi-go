package xivapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

// DefaultResultsPerPage is the default amount of results to return per page
const DefaultResultsPerPage int = 100

// SearchParams is the struct that gets encoded to the list of uri params
type SearchParams struct {
	GlobalQueryParameters

	Indexes      SearchIndexes `url:"indexes"`
	Query        string        `url:"string"`
	Algorithm    SearchAlgo    `url:"string_algo,omitempty"`
	SearchColumn string        `url:"string_column,omitempty"`
	Page         int           `url:"page,omitempty"`
	SortColumn   string        `url:"sort_field,omitempty"`
	SortOrder    SortOrder     `url:"sort_order,omitempty"`
	Limit        int           `url:"limit"`
	Filters      SearchFilters `url:"filters"`
}

// SearchResult is the result returned from XIVAPI
type SearchResult struct {
	Pagination   Pagination
	Results      []SearchResultEntry `json:"results"`
	Milliseconds int64               `json:"SpeedMs"`
}

func (sr SearchResult) String() string {
	return fmt.Sprintf("SearchResult{ %v %v }", sr.Pagination, sr.Results)
}

// Duration converts the Milliseconds into a time.Duraction
func (sr *SearchResult) Duration() time.Duration {
	return time.Duration(sr.Milliseconds) * time.Millisecond
}

// SearchResultEntry is the result of a search result.
// It holds the 2 common fields to all results and allows converting into the correct type
type SearchResultEntry struct {
	*SearchResultCommon
	raw json.RawMessage
}

// UnmarshalJSON is called by encoding/json for a custom way to unmarshal the json object
func (e *SearchResultEntry) UnmarshalJSON(bs []byte) error {
	var decoded SearchResultCommon
	if err := json.Unmarshal(bs, &decoded); err != nil {
		return err
	}

	e.SearchResultCommon = &decoded
	e.raw = bs

	return nil
}

func (e SearchResultEntry) String() string {
	return fmt.Sprintf("%v { %v }", e.Type, e.Name)
}

// Search prepares a simple search request
// This uses SearchRaw under the hood
func (c *XIVAPI) Search(query, searchColumn, sortColumn string, sortOrder SortOrder, page int, limit int, filters SearchFilters, indexes ...SearchIndex) (*SearchResult, error) {
	params := SearchParams{
		Query:        query,
		SearchColumn: searchColumn,
		SortColumn:   sortColumn,
		SortOrder:    sortOrder,
		Limit:        limit,
		Page:         page,
		Filters:      filters,
		Indexes:      SearchIndexes(indexes),
	}

	if limit <= 0 {
		params.Limit = DefaultResultsPerPage
	}

	return c.SearchRaw(params)
}

// SearchRaw accepts a raw SearchParams argument and executes the search
func (c *XIVAPI) SearchRaw(params SearchParams) (*SearchResult, error) {
	// xivapi uses 1-based pagination
	// lib uses 0-based indexing
	params.Page++

	qs, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	uri, _ := url.Parse(fmt.Sprintf("%v%v", BaseURL, "/Search"))
	uri.RawQuery = qs.Encode()

	r := new(SearchResult)
	if err := c.RequestJSON(MethodGet, uri, nil, r); err != nil {
		return nil, err
	}

	return r, nil
}
