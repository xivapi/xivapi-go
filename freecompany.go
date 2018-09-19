package xivapi

import (
	"fmt"
	"net/url"
)

// FreeCompanyState : Shows current status of FreeCompany on XIVAPI
type FreeCompanyState int

const (
	// StateNone : Content is not on XIVAPI and will not be added via this request
	StateNone FreeCompanyState = iota
	// StateAdding : Content does not exist on the API and needs adding.
	StateAdding
	// StateCached : Content exists in the system and you're being provided a cached response.
	StateCached
	// StateNotFound : Content does not exist on The Lodestone.
	StateNotFound
	// StateBlacklist : Content has been Blacklisted. No data can be obtained via the API for any application
	StateBlacklist
	// StatePrivate : Content is private on lodestone, ask the owner to make the content public and then try again!
	StatePrivate
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

// FreeCompany requests information about the provided free company ID
// This call isn't implemented by XIVAPI yet, so this is pretty much a no-op call.
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
func (c *XIVAPI) FreeCompanySearch(name, server, page string) (*SearchFreeCompanyResult, error) {
	if name == "" {
		return nil, ErrMissingKey
	}
	uri, _ := url.Parse(fmt.Sprintf("%v%v%v", BaseURL, freeCompanyEndpoint, "Search"))

	values := uri.Query()
	values.Add("name", name)

	if server != "" {
		values.Add("server", server)
	}
	if page != "" {
		values.Add("page", page)
	}
	uri.RawQuery = values.Encode()

	r := new(SearchFreeCompanyResult)
	if err := c.RequestJSON(MethodGet, uri, nil, r); err != nil {
		return nil, err
	}

	return r, nil
}
