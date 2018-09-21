package xivapi

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

const freeCompanyEndpoint = "/FreeCompany/"

type SearchFreeCompanyResult struct {
	Pagination Pagination                    `json:"pagination"`
	Results    []SearchFreeCompanyResultItem `json:"results"`
}

type SearchFreeCompanyResultItem struct {
	Crest  []string `json:"Crest"`
	ID     string   `json:"ID"`
	Name   string   `json:"Name"`
	Server string   `json:"Server"`
}

//PlayerContentQueryOptions provide search parameters for all /*/Search endpoints (PvPTeam, Character, FreeCompany, LinkShell)
type PlayerContentQueryOptions struct {
	Name   string `url:"name,omitempty"`
	Server string `url:"server,omitempty"`
	Page   int    `url:"page,omitempty"`
}

// FreeCompany requests information about the provided free company ID
func (c *XIVAPI) FreeCompany(id uint64, files ...FileType) (interface{}, error) {
	uri, _ := url.Parse(fmt.Sprintf("%v%v%v", BaseURL, freeCompanyEndpoint, id))

	// force record info
	if len(files) == 0 {
		files = append(files, FileRecord)
	}

	filesString := ""
	for _, f := range files {
		filesString += string(f) + ","
	}

	values := uri.Query()
	values.Set("files", filesString[:len(filesString)-1])
	uri.RawQuery = values.Encode()

	return nil, ErrNotImplemented
}

// FreeCompanySearch searches for a FreeCompany on XIVAPI provided a free company name
func (c *XIVAPI) FreeCompanySearch(name, server string, page int) (*SearchFreeCompanyResult, error) {
	if name == "" {
		return nil, ErrMissingKey
	}
	uri, _ := url.Parse(fmt.Sprintf("%v%v%v", BaseURL, freeCompanyEndpoint, "Search"))

	opt := &PlayerContentQueryOptions{name, server, page}
	v, _ := query.Values(opt)
	uri.RawQuery = v.Encode()

	r := new(SearchFreeCompanyResult)
	if err := c.RequestJSON(MethodGet, uri, nil, r); err != nil {
		return nil, err
	}

	return r, nil
}
