package xivapi

// SearchAlgo specifies the available algorithms for searching
type SearchAlgo string

const (
	// AlgoWildcard is a very basic wild card, for example: ard would match: b-ard-ing or h-ard etc.
	AlgoWildcard SearchAlgo = "wildcard"

	// undocumented, thus private
	algoMultiMatch        = "multi_match"
	algoQueryString       = "query_string"
	algoTerm              = "term"
	algoMatchPhrasePrefix = "match_phrase_prefix"
	algoFuzzy             = "fuzzy"
)

// SortOrder defines the sort order for the result list
type SortOrder string

const (
	// OrderAscending sorts ascending
	OrderAscending SortOrder = "asc"
	// OrderDescending sorts descending
	OrderDescending = "desc"
)
