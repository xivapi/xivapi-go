package xivapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/google/go-querystring/query"
)

// DefaultResultsPerPage is the default amount of results to return per page
const DefaultResultsPerPage int = 100

// SearchParams is the struct that gets encoded to the list of uri params
type SearchParams struct {
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
	Milliseconds int64               `json:"ms"`
	Pagination   Pagination          `json:"pagination"`
	Results      []SearchResultEntry `json:"results"`
}

func (sr SearchResult) String() string {
	return fmt.Sprintf("SearchResult{ %v %v }", sr.Pagination, sr.Results)
}

// Duration converts the Milliseconds into a time.Duraction
func (r *SearchResult) Duration() time.Duration {
	return time.Duration(r.Milliseconds) * time.Millisecond
}

type SearchResultEntry struct {
	Type SearchIndex `json:"_"`
	Name string      `json:"name"`

	rawMap map[string]interface{} `json:"-"`
}

func (e *SearchResultEntry) UnmarshalJSON(bs []byte) error {
	// encoding/json converts snake_case to PascalCase
	if err := json.Unmarshal(bs, &e.rawMap); err != nil {
		return err
	}

	if _, ok := e.rawMap["_"]; !ok {
		return ErrMissingKey
	}

	if _, ok := e.rawMap["Name"]; !ok {
		return ErrMissingKey
	}

	if v, ok := e.rawMap["_"].(string); ok {
		e.Type = SearchIndex(v)
	} else {
		return ErrUnexpectedType
	}

	if v, ok := e.rawMap["Name"].(string); ok {
		e.Name = v
	} else {
		return ErrUnexpectedType
	}

	return nil
}

func (e *SearchResultEntry) GetCompanion() (*SearchResultCompanion, error) {
	entity := new(SearchResultCompanion)
	if err := mapstructure.Decode(e.rawMap, entity); err != nil {
		return nil, err
	}
	return entity, nil
}

func (e SearchResultEntry) String() string {
	return fmt.Sprintf("%v { %v }", e.Type, e.Name)
}

// Search prepares a simple search request
// This uses SearchRaw under the hood
func (c *XIVAPI) Search(query, searchColumn, sortColumn string, sortOrder SortOrder, page int, limit int, filters SearchFilters, indexes ...SearchIndex) (*SearchResult, error) {
	if limit <= 0 {
		limit = DefaultResultsPerPage
	}

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

	return c.SearchRaw(params)
}

// SearchRaw accepts a raw SearchParams argument and executes the search
func (c *XIVAPI) SearchRaw(params SearchParams) (*SearchResult, error) {
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
