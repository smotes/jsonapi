package jsonapi

// Error represents a JSON API error object.
//
// http://jsonapi.org/format/#error-objects
type Error struct {
	ID     string                 `json:"id,omitempty"`
	Status string                 `json:"status,omitempty"`
	Code   string                 `json:"code,omitempty"`
	Title  string                 `json:"title,omitempty"`
	Detail string                 `json:"detail,omitempty"`
	Links  Links                  `json:"links,omitempty"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
}
