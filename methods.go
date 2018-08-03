package xivapi

// Methods enumerates every supported HTTP Method
type Methods string

// Collection of all supported HTTP Methods
const (
	MethodGet    Methods = "GET"
	MethodPost           = "POST"
	MethodPut            = "PUT"
	MethodDelete         = "DELETE"
)
