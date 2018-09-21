package xivapi

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const freeCompanyEndpoint = "/FreeCompany/"

// FreeCompanySearch searches for a FreeCompany on XIVAPI provided a free company name
func (c *XIVAPI) FreeCompanySearch(name, server string, page int) (*SearchFreeCompanyResult, error) {
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

// FreeCompany requests information about the provided free company ID
func (c *XIVAPI) FreeCompany(id string, isMembersIncluded bool, columns ...string) (*FreeCompanyResult, error) {
	uri, _ := url.Parse(fmt.Sprintf("%v%v%v", BaseURL, freeCompanyEndpoint, id))
	q := buildFreeCompanyQuery(isMembersIncluded, columns)
	v, _ := query.Values(q)
	uri.RawQuery = v.Encode()

	r := new(FreeCompanyResult)
	if err := c.RequestJSON(MethodGet, uri, nil, r); err != nil {
		return nil, err
	}
	return r, nil
}

func buildFreeCompanyQuery(m bool, c []string) *FreeCompanyQuery {
	r := new(FreeCompanyQuery)
	r.Columns = strings.Join(c, ",")

	if m {
		r.Data = "FCM"
	}

	return r
}
