package jsonapi_test

import (
	"testing"

	"github.com/smotes/jsonapi"
)

func TestFromResource_WhenNoAdapter(t *testing.T) {
	s := struct{}{}
	r := &jsonapi.Resource{}

	err := jsonapi.FromResource(s, r, true)
	if err == nil {
		t.Error("FromResource() should return error when adapter does not satisfy identity write interface")
	}
}

// id

type badIDWriteAdapter struct{}

func (a *badIDWriteAdapter) SetID(id string) error {
	return testErr
}

func (a *badIDWriteAdapter) SetType(typ string) error {
	return nil
}

func TestFromResource_WhenSetIDError(t *testing.T) {
	testFromResourceForError(t, &badIDWriteAdapter{}, "SetID", testErr)
}

// type

type badTypeWriteAdapter struct{}

func (a *badTypeWriteAdapter) SetID(id string) error {
	return nil
}

func (a *badTypeWriteAdapter) SetType(typ string) error {
	return testErr
}

func TestFromResource_WhenSetTypeError(t *testing.T) {
	testFromResourceForError(t, &badTypeWriteAdapter{}, "SetType", testErr)
}

// attributes

type badAttributesWriteAdapter struct{}

func (a *badAttributesWriteAdapter) SetID(id string) error {
	return nil
}

func (a *badAttributesWriteAdapter) SetType(typ string) error {
	return nil
}

func (a *badAttributesWriteAdapter) SetAttributes(v map[string]interface{}) error {
	return testErr
}

func TestFromResource_WhenSetAttributesError(t *testing.T) {
	testFromResourceForError(t, &badAttributesWriteAdapter{}, "SetAttributes", testErr)
}

// relationships

type badRelationshipsWriteAdapter struct{}

func (a *badRelationshipsWriteAdapter) SetID(id string) error {
	return nil
}

func (a *badRelationshipsWriteAdapter) SetType(typ string) error {
	return nil
}

func (a *badRelationshipsWriteAdapter) SetRelationships(rs jsonapi.Relationships) error {
	return testErr
}

func TestFromResource_WhenSetRelationshipsError(t *testing.T) {
	testFromResourceForError(t, &badRelationshipsWriteAdapter{}, "SetRelationships", testErr)
}

type notFullWriteAdapter struct{}

func (a *notFullWriteAdapter) SetID(id string) error {
	return nil
}

func (a *notFullWriteAdapter) SetType(typ string) error {
	return nil
}

func (a *notFullWriteAdapter) SetAttributes(attrs map[string]interface{}) error {
	return testErr
}

func (a *notFullWriteAdapter) SetRelationships(rels jsonapi.Relationships) error {
	return testErr
}

func TestFromResource_WhenNotFull(t *testing.T) {
	adapter := notFullWriteAdapter{}
	r := jsonapi.Resource{}
	if err := jsonapi.FromResource(&adapter, &r, false); err != nil {
		t.Error("unexpected error when calling FromResource with full=false, should not set attributes or relationships")
	}
}

// helpers

func testFromResourceForError(t *testing.T, adapter interface{}, method string, expected error) {
	r := jsonapi.Resource{ID: "test"}
	if actual := jsonapi.FromResource(adapter, &r, true); actual != expected {
		t.Errorf("expected FromResource to return error from adapter.%s, expected: %v, got: %v",
			method, expected, actual)
	}

}
