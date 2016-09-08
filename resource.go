package jsonapi

import "fmt"

// Resource represents a JSON API resource identifier or resource object.
// Each valid JSON API resource object must contain at least the "id" and "type" keys,
// and may contain the optional "attributes", "relationships", "links" and "meta" keys.
//
// Note that any of the optional keys will be omitted if a value is not provided.
//
// For more information, see the specification at:
//
// http://jsonapi.org/format/#document-resource-object
//
// http://jsonapi.org/format/#document-resource-identifier-objects
type Resource struct {
	ID            string                 `json:"id"`
	Type          string                 `json:"type"`
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
	Relationships Relationships          `json:"relationships,omitempty"`
	Links         Links                  `json:"links,omitempty"`
	Meta          map[string]interface{} `json:"meta,omitempty"`
}

// ToResource uses the adapter implementation v to return the corresponding Resource.
//
// If full is true, ToResource will attempt to set the resource object's optional members
// including Attributes, Links and Meta. If full is false, the resultant resource object
// will only include the ID and Type members, allowing it to be used as a resource identifier
// object.
//
// Several adapter interfaces are used to convert a custom type to a Resource; one is required
// and the rest are optional.
//
//	type identityReadAdapter interface {
//		GetID() (string, error)
//		GetType() (string, error)
//	}
//
// The identity read adapter is used to populate the ID and Type fields on the resultant Resource.
// It is required that all custom types implement this interface or ToResource will return an error.
//
// 	type attributesReadAdapter interface {
// 		GetAttributes() (map[string]interface{}, error)
// 	}
//
// The attributes read adapter is used to populate the Attributes field on the resultant Resource.
// It is optional and will be ignored if not implemented, or if it returns nil.
//
// 	type relationshipsReadAdapter interface {
// 		GetRelationships() (jsonapi.Relationships, error)
// 	}
//
// The relationships read adapter is used to populate the Relationships field on the resultant Resource.
// It is optional and will be ignored if not implemented, or if it returns nil.
//
// 	type linksReadAdapter interface {
// 		GetLinks() (jsonapi.Links, error)
// 	}
//
// The links read adapter is used to populate the Links field on the resultant Resource.
// It is optional and will be ignored if not implemented, or if it returns nil.
//
// 	type metaReadAdapter interface {
// 		GetMeta() (map[string]interface{}, error)
// 	}
//
// The meta read adapter is used to populate the Meta field on the resultant Resource.
// It is optional and will be ignored if not implemented, or if it returns nil.
func ToResource(v interface{}, full bool) (*Resource, error) {
	// each resource must have at least the "id" and "type" members
	adapter, ok := v.(identityReadAdapter)
	if !ok {
		return nil, errResourceIdentity
	}

	var (
		id, typ string
		err     error
	)

	id, err = adapter.GetID()
	if err != nil {
		return nil, err
	}
	typ, err = adapter.GetType()
	if err != nil {
		return nil, err
	}

	r := &Resource{
		ID:   id,
		Type: typ,
	}

	if !full {
		return r, nil
	}

	// add optional members to resource if their respective read adapters are satisfied
	if v, ok := v.(attributesReadAdapter); ok {
		r.Attributes, err = v.GetAttributes()
		if err != nil {
			return nil, err
		}
	}
	if v, ok := v.(relationshipsReadAdapter); ok {
		r.Relationships, err = v.GetRelationships()
		if err != nil {
			return nil, err
		}
	}
	if v, ok := v.(linksReadAdapter); ok {
		r.Links, err = v.GetLinks()
		if err != nil {
			return nil, err
		}
	}
	if v, ok := v.(metaReadAdapter); ok {
		r.Meta, err = v.GetMeta()
		if err != nil {
			return nil, err
		}
	}

	return r, nil
}

// FromResource uses the adapter implementation v to set the values from the corresponding Resource r.
//
// Several adapter interfaces are used to convert a Resource to a custom type; one is required
// and the rest are optional.
//
//	type identityWriteAdapter interface {
//		SetID(string) error
//		SetType(string) error
//	}
//
// The identity write adapter is used to populate the ID field on custom type from the Resource.
// It is required that all custom types implement this interface or FromResource will return an error.
//
// 	type attributesWriteAdapter interface {
// 		SetAttributes(map[string]interface{}) error
// 	}
//
// The attributes write adapter is used to populate fields on the custom type from the Attributes on the Resource.
// It is optional and will be ignored if not implemented, or if it returns nil.
//
// 	type relationshipsWriteAdapter interface {
// 		SetRelationships(jsonapi.Relationships) error
// 	}
//
// The relationships write adapter is used to populate fields on the custom type from the Relationships on the Resource.
// It is optional and will be ignored if not implemented, or if it returns nil.
//
// Note the lack of a linksWriteAdapter and metaWriteAdapter, or a SetType method on the identityWriterAdapter. As per
// the JSON API specification for client interaction with an API (including reading, creating, updating or deleting any
// resources or relationships), these adapters should be unnecessary.
func FromResource(adapter interface{}, r *Resource, full bool) error {
	v, ok := adapter.(identityWriteAdapter)
	if !ok {
		return errResourceIdentity
	}

	// skip setting ID for edge case when creating new resource
	if len(r.ID) > 0 {
		if err := v.SetID(r.ID); err != nil {
			return err
		}
	}

	if err := v.SetType(r.Type); err != nil {
		return err
	}

	if !full {
		return nil
	}

	if v, ok := adapter.(attributesWriteAdapter); ok {
		if err := v.SetAttributes(r.Attributes); err != nil {
			return err
		}
	}
	if v, ok := adapter.(relationshipsWriteAdapter); ok {
		if err := v.SetRelationships(r.Relationships); err != nil {
			return err
		}
	}

	return nil
}

// identity adapters

type identityReadAdapter interface {
	GetID() (string, error)
	GetType() (string, error)
}

type identityWriteAdapter interface {
	SetID(string) error
	SetType(string) error
}

// attributes adapters

type attributesReadAdapter interface {
	GetAttributes() (map[string]interface{}, error)
}

type attributesWriteAdapter interface {
	SetAttributes(map[string]interface{}) error
}

// relationships adapters

type relationshipsReadAdapter interface {
	GetRelationships() (Relationships, error)
}

type relationshipsWriteAdapter interface {
	SetRelationships(Relationships) error
}

// links adapters

type linksReadAdapter interface {
	GetLinks() (Links, error)
}

// meta adapters

type metaReadAdapter interface {
	GetMeta() (map[string]interface{}, error)
}

// errors

var errResourceIdentity = fmt.Errorf("%s: invalid resource identity", packageName)
