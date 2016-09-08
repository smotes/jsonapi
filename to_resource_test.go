package jsonapi_test

import (
	"testing"

	"github.com/smotes/jsonapi"
)

func TestToResource_WhenNoAdapter(t *testing.T) {
	s := struct{}{}

	r, err := jsonapi.ToResource(s, false)
	if r != nil || err == nil {
		t.Error("ToResource() should return nil/error when adapter does not satisfy identity read interface")
	}
}

// id

type badIDReadAdapter struct{}

func (a *badIDReadAdapter) GetID() (string, error) {
	return "", testErr
}

func (a *badIDReadAdapter) GetType() (string, error) {
	return "tests", nil
}

func TestToResource_WhenGetIDError(t *testing.T) {
	adapter := badIDReadAdapter{}
	r, err := jsonapi.ToResource(&adapter, false)
	if r != nil || err != testErr {
		t.Error("ToResource should return nil/error from adapter.GetID when it returns an error")
	}
}

// type

type badTypeReadAdapter struct{}

func (a *badTypeReadAdapter) GetID() (string, error) {
	return "1", nil
}

func (a *badTypeReadAdapter) GetType() (string, error) {
	return "", testErr
}

func TestToResource_WhenGetTypeError(t *testing.T) {
	adapter := badTypeReadAdapter{}
	r, err := jsonapi.ToResource(&adapter, false)
	if r != nil || err != testErr {
		t.Error("ToResource should return nil/error from adapter.GetType when it returns an error")
	}
}

// attributes

type badAttributesReadAdapter struct{}

func (a *badAttributesReadAdapter) GetID() (string, error) {
	return "1", nil
}

func (a *badAttributesReadAdapter) GetType() (string, error) {
	return "tests", nil
}

func (a *badAttributesReadAdapter) GetAttributes() (map[string]interface{}, error) {
	return nil, testErr
}

func TestToResource_WhenGetAttributesError(t *testing.T) {
	adapter := badAttributesReadAdapter{}
	r, err := jsonapi.ToResource(&adapter, true)
	if r != nil || err != testErr {
		t.Error("ToResource should return nil/error from adapter.GetAttributes when it returns an error")
	}
}

// relationships

type badRelationshipsReadAdapter struct{}

func (a *badRelationshipsReadAdapter) GetID() (string, error) {
	return "1", nil
}

func (a *badRelationshipsReadAdapter) GetType() (string, error) {
	return "tests", nil
}

func (a *badRelationshipsReadAdapter) GetRelationships() (jsonapi.Relationships, error) {
	return nil, testErr
}

func TestToResource_WhenGetRelationshipsError(t *testing.T) {
	adapter := badRelationshipsReadAdapter{}
	r, err := jsonapi.ToResource(&adapter, true)
	if r != nil || err != testErr {
		t.Error("ToResource should return nil/error from adapter.GetRelationships when it returns an error")
	}
}

// links

type badLinksReadAdapter struct{}

func (a *badLinksReadAdapter) GetID() (string, error) {
	return "1", nil
}

func (a *badLinksReadAdapter) GetType() (string, error) {
	return "tests", nil
}

func (a *badLinksReadAdapter) GetLinks() (jsonapi.Links, error) {
	return nil, testErr
}

func TestToResource_WhenGetLinksError(t *testing.T) {
	adapter := badLinksReadAdapter{}
	r, err := jsonapi.ToResource(&adapter, true)
	if r != nil || err != testErr {
		t.Error("ToResource should return nil/error from adapter.GetLinks when it returns an error")
	}
}

// meta

type badMetaReadAdapter struct{}

func (a *badMetaReadAdapter) GetID() (string, error) {
	return "1", nil
}

func (a *badMetaReadAdapter) GetType() (string, error) {
	return "tests", nil
}

func (a *badMetaReadAdapter) GetMeta() (map[string]interface{}, error) {
	return nil, testErr
}

func TestToResource_WhenGetMetaError(t *testing.T) {
	adapter := badMetaReadAdapter{}
	r, err := jsonapi.ToResource(&adapter, true)
	if r != nil || err != testErr {
		t.Error("ToResource should return nil/error from adapter.GetMeta when it returns an error")
	}
}
