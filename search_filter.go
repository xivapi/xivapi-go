package xivapi

import (
	"fmt"
	"net/url"
	"strings"
)

// FilterOperator defines the different filter operators for extended filtering
type FilterOperator string

// The different filters
const (
	FilterEquals         FilterOperator = "="
	FilterGreaterOrEqual                = ">"
	FilterLessOrEqual                   = "<"
)

// SearchFilter is a singular filter
type SearchFilter struct {
	Field    string
	Operator FilterOperator
	Value    string
}

func (sf SearchFilter) String() string {
	return fmt.Sprintf("%v%v%v", sf.Field, string(sf.Operator), sf.Value)
}

// SearchFilters is a list of ilters
type SearchFilters []SearchFilter

// AddFilter is a convinience wrapper for SearchFilters to create and add a filter in a single line
func (sf SearchFilters) AddFilter(field string, operator FilterOperator, value string) SearchFilters {
	return append(sf, SearchFilter{field, operator, value})
}

func (sf SearchFilters) String() string {
	filtersMap := make(map[string]struct{})
	for _, filter := range sf {
		filtersMap[filter.String()] = struct{}{}
	}

	filters := make([]string, len(filtersMap))
	i := 0
	for index := range filtersMap {
		filters[i] = index
		i++
	}

	return strings.Join(filters, ",")
}

// EncodeValues encodes the list into a custom list of values
// This is a custom Encoder for go-querystring https://godoc.org/github.com/google/go-querystring/query
func (sf SearchFilters) EncodeValues(key string, v *url.Values) error {
	v.Set(key, sf.String())
	return nil
}
