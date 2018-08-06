package xivapi

import (
	"fmt"
	"net/url"
)

func (c *XIVAPI) Linkshell(id uint64, files ...FileType) (interface{}, error) {
	uri, _ := url.Parse(fmt.Sprintf("%v%v%v", BaseURL, "/Linkshell/", id))

	filesString := ""
	for _, f := range files {
		filesString += string(f) + ","
	}

	values := uri.Query()
	values.Set("files", filesString[:len(filesString)-1])
	uri.RawQuery = values.Encode()

	return nil, ErrNotImplemented
}
