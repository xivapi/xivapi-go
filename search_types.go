package xivapi

// SearchAlgo specifies the available algorithms for searching
type SearchAlgo string

// The different search algorithm.
// See https://xivapi.com/docs/Search#string_algo for details
const (
	AlgoWildcard          SearchAlgo = "wildcard"
	AlgoWildcardPlus                 = "wildcard_plus"
	AlgoMatchPhrasePrefix            = "match_phrase_prefix"
	AlgoMultiMatch                   = "multi_match"
	QueryString                      = "query_string"
	Term                             = "term"
	Fuzzy                            = "fuzzy"
)

// SortOrder defines the sort order for the result list
type SortOrder string

const (
	// OrderAscending sorts ascending
	OrderAscending SortOrder = "asc"
	// OrderDescending sorts descending
	OrderDescending = "desc"
)
