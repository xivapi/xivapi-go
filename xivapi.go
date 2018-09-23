package xivapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

// BaseURL is the root URL of every HTTP request
const BaseURL string = "https://xivapi.com"

const (
	defaultTimeout = 30 * time.Second
)

// XIVAPI is the client to the xivapi.com website
type XIVAPI struct {
	WebClient *http.Client

	PreRequest  func(*http.Request) error
	PostRequest func(*http.Response) error

	key string
}

// New creates a new XIVAPI client
func New(key string) *XIVAPI {
	c := new(XIVAPI)
	c.key = key

	c.WebClient = &http.Client{
		Timeout: defaultTimeout,
	}

	return c
}

// Request issues an HTTP request to the xivapi.com website and returns the raw http body as a reader
func (c *XIVAPI) Request(method Methods, uri *url.URL, body io.ReadCloser) (io.ReadCloser, error) {
	query := uri.Query()
	query.Set("key", c.key)
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest(string(method), uri.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Body = body
	}

	if c.PreRequest != nil {
		if err := c.PreRequest(req); err != nil {
			return nil, err
		}
	}

	req.Header.Add("User-Agent", fmt.Sprintf("gitlab.com-paars-xiv-xivapi/%v %v (%v)", Version, runtime.GOOS, runtime.GOARCH))
	res, err := c.WebClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, NewStatusCodeError(res.StatusCode)
	}

	if c.PostRequest != nil {
		if err := c.PostRequest(res); err != nil {
			return nil, err
		}
	}

	return res.Body, nil
}

// RequestJSON calls Request to get the io.ReadCloser and then converts the body to the specified JSON object
func (c *XIVAPI) RequestJSON(method Methods, uri *url.URL, body io.ReadCloser, jsonStruct interface{}) error {
	responseBody, err := c.Request(method, uri, body)
	if err != nil {
		return err
	}

	defer responseBody.Close()
	json.NewDecoder(responseBody).Decode(jsonStruct)
	return nil
}
