package jsonapi

// Links represents a JSON API links object, which can include JSON API link objects or strings representing links.
//
// http://jsonapi.org/format/#document-links
type Links map[string]interface{}

// Link represents a JSON API link object, which includes the required "href" key and optional "meta" key.
//
// http://jsonapi.org/format/#document-links
type Link struct {
	Href string                 `json:"href"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

// AddString adds the string key, value pair to the meta object.
// It overwrites any existing values associated with key.
func (ls Links) AddString(key, v string) {
	if ls == nil {
		return
	}
	ls[key] = v
}

// Add adds the object key, value pair to the meta object.
// It overwrites any existing values associated with key.
func (ls Links) Add(key string, v *Link) {
	if ls == nil {
		return
	}
	ls[key] = v
}

// GetString returns the string value associated with the given key and an existence check.
// Returns an empty string/false if the value associated with the key does not exist or the value is not of type string.
func (ls Links) GetString(key string) (string, bool) {
	v, ok := ls.getKey(key)
	if !ok {
		return "", false
	}

	switch typ := v.(type) {
	case string:
		return typ, true
	default:
		return "", false
	}
}

// Get returns the object value associated with the given key and an existence check.
// Returns a nil/false if the value associated with the key does not exist or does not reference a Link struct.
func (ls Links) Get(key string) (*Link, bool) {
	v, ok := ls.getKey(key)
	if !ok {
		return nil, false
	}

	switch typ := v.(type) {
	case map[string]interface{}:
		l := Link{}
		if href, ok := typ["href"]; ok {
			if href, ok := href.(string); ok {
				l.Href = href
			}
		}
		if meta, ok := typ["meta"]; ok {
			if meta, ok := meta.(map[string]interface{}); ok {
				l.Meta = meta
			}
		}
		return &l, true

	case *Link:
		return typ, true

	default:
		return nil, false
	}
}

// Delete deletes the value associated with the given key.
func (ls Links) Delete(key string) {
	if ls == nil {
		return
	}
	delete(ls, key)
}

func (ls Links) getKey(key string) (v interface{}, ok bool) {
	if ls == nil {
		return nil, false
	}
	v, ok = ls[key]
	return
}
