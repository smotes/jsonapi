package jsonapi

import "encoding/json"

// Document represents top-level document at the root of any JSON API request/response containing data.
//
// http://jsonapi.org/format/#document-top-level
type Document struct {
	Data     json.RawMessage        `json:"data,omitempty"`
	Errors   []Error                `json:"errors,omitempty"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
	Info     *Info                  `json:"jsonapi,omitempty"`
	Links    Links                  `json:"links,omitempty"`
	Included []Resource             `json:"included,omitempty"`
}

// Info represents a JSON API Object, used as the "jsonapi" member in the top-level document, which provides
// information about its implementation.
//
// http://jsonapi.org/format/#document-jsonapi-object
type Info struct {
	Version string                 `json:"version"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}
