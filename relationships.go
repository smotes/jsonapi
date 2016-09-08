package jsonapi

import "encoding/json"

// Relationships represents a JSON API relationships object
//
// http://jsonapi.org/format/#document-resource-object-relationships
type Relationships map[string]*Relationship

// Relationship represents a member of the JSON API relationships object.
type Relationship struct {
	Links Links                  `json:"links,omitempty"`
	Data  json.RawMessage        `json:"data,omitempty"`
	Meta  map[string]interface{} `json:"meta,omitempty"`
}

// Add adds the key, value pair to the meta object.
// It overwrites any existing values associated with key.
func (rs Relationships) Add(key string, r *Relationship) {
	if rs == nil {
		return
	}
	rs[key] = r
}

// Get returns the value associated with the given key and an existence check.
func (rs Relationships) Get(key string) (*Relationship, bool) {
	if rs == nil {
		return nil, false
	}
	r, ok := rs[key]
	return r, ok
}

// Delete deletes the value associated with the given key.
func (rs Relationships) Delete(key string) {
	if rs == nil {
		return
	}
	delete(rs, key)
}
