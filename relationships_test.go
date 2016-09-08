package jsonapi_test

import (
	"testing"

	"github.com/smotes/jsonapi"
)

func TestRelationships_Add_WhenNil(t *testing.T) {
	defer catchPanic(t, "Relationships", "Add")
	var rs jsonapi.Relationships
	rs.Add("foo", &jsonapi.Relationship{})
}

func TestRelationships_Add(t *testing.T) {
	var (
		rs       = jsonapi.Relationships{}
		key      = "foo"
		expected = &jsonapi.Relationship{}
	)

	rs.Add(key, expected)
	if actual, ok := rs[key]; actual == nil || !ok {
		t.Error("underlying map in Relationships should have value associated with key after Relationships.Add")
	} else if actual != expected {
		t.Errorf("unexpected value after Relationships.Add, expected: %s, actual: %s",
			expected, actual)
	}
}

func TestRelationships_Get_WhenNil(t *testing.T) {
	defer catchPanic(t, "Relationships", "Get")
	var rs jsonapi.Relationships
	rs.Get("foo")
}

func TestRelationships_Get(t *testing.T) {
	var (
		rs       = jsonapi.Relationships{}
		key      = "foo"
		expected = &jsonapi.Relationship{}
	)

	if actual, ok := rs.Get(key); actual != nil || ok {
		t.Error("Relationships.Get should return nil/false when no value is associated with the key")
	}

	rs.Add(key, expected)
	if actual, ok := rs.Get(key); !ok {
		t.Error("expected Relationships.Get to return true existence check after Relationships.Add for the given key")
	} else if actual != expected {
		t.Errorf("unexpected value after Relationships.Get, expected: %s, actual: %s",
			expected, actual)
	}
}

func TestRelationships_Delete_WhenNil(t *testing.T) {
	defer catchPanic(t, "Relationships", "Delete")
	var rs jsonapi.Relationships
	rs.Delete("foo")
}

func TestRelationships_Delete(t *testing.T) {
	var (
		rs  = jsonapi.Relationships{}
		key = "foo"
	)

	rs.Add(key, &jsonapi.Relationship{})
	rs.Delete(key)
	if actual, ok := rs.Get(key); actual != nil || ok {
		t.Error("Relationships.Get should return nil/false after calling Relationships.Delete for the given key")
	}
}
