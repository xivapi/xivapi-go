package xivapi

import (
	"fmt"
	"net/url"
)

// FreeCompany requests information about the provided free company ID
// This call isn't implemented by XIVAPI yet, so this is pretty much a no-op call.
func (c *XIVAPI) FreeCompany(id uint64, files ...FileType) (interface{}, error) {
	uri, _ := url.Parse(fmt.Sprintf("%v%v%v", BaseURL, "/FreeCompany/", id))

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
