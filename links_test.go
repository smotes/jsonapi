package jsonapi_test

import (
	"reflect"
	"testing"

	"github.com/smotes/jsonapi"
)

func TestLinks_AddString_WhenNil(t *testing.T) {
	var ls jsonapi.Links = nil
	defer catchPanic(t, "Links", "AddString")
	ls.AddString("foo", "")
}

func TestLinks_AddString(t *testing.T) {
	var (
		ls       = jsonapi.Links{}
		key      = "foo"
		expected = "bar"
	)

	ls.AddString(key, expected)
	if v, ok := ls[key]; !ok {
		t.Error("underlying map in Links should have value associated with key after Links.AddString")

		if actual, ok := v.(string); !ok {
			t.Error("unexpected value type after Links.AddString, expected string")
		} else if actual != expected {
			t.Errorf("unexpected value after Links.AddString, expected: %s, actual: %s", expected, actual)
		}
	}
}

func TestLinks_Add_WhenNil(t *testing.T) {
	var ls jsonapi.Links = nil
	defer catchPanic(t, "Links", "Add")
	ls.Add("foo", &jsonapi.Link{})
}

func TestLinks_Add(t *testing.T) {
	var (
		ls       = jsonapi.Links{}
		key      = "foo"
		expected = &jsonapi.Link{}
	)

	ls.Add(key, expected)
	if v, ok := ls[key]; !ok {
		t.Error("underlying map in Links should have value associated with key after Links.Add")
	} else if actual, ok := v.(*jsonapi.Link); !ok {
		t.Error("unexpected value type after Links.Add, expected *Link")
	} else if actual != expected {
		t.Errorf("unexpected value after Links.Add, expected: %s, actual: %s", expected, actual)
	}
}

func TestLinks_GetString_WhenNil(t *testing.T) {
	var ls jsonapi.Links = nil
	defer catchPanic(t, "Links", "GetString")
	ls.GetString("foo")
}

func TestLinks_GetString(t *testing.T) {
	var (
		ls       = jsonapi.Links{}
		key      = "foo"
		expected = "bar"
	)

	ls = jsonapi.Links{}
	if actual, ok := ls.GetString(key); len(actual) > 0 || ok {
		t.Error("expected Links.GetString to return empty string/false before Links.AddString for the given key")
	}

	ls.AddString(key, expected)
	if actual, ok := ls.GetString(key); !ok {
		t.Error("expected Links.GetString to return true existence check after Links.AddString for the given key")
	} else if actual != expected {
		t.Errorf("unexpected value after Links.GetString, expected: %s, actual: %s",
			expected, actual)
	}
}

func TestLinks_GetString_WhenWrongType(t *testing.T) {
	var (
		ls  = jsonapi.Links{}
		key = "foo"
	)

	ls.Add(key, &jsonapi.Link{})
	_, ok := ls.GetString(key)
	if ok {
		t.Error("Links.GetString should return empty string/false when the value associated with the key is not of type string")
	}
}

func TestLinks_Get_WhenNil(t *testing.T) {
	var ls jsonapi.Links = nil
	defer catchPanic(t, "Links", "Get")
	ls.Get("foo")
}

func TestLinks_Get(t *testing.T) {
	var (
		ls       = jsonapi.Links{}
		key      = "foo"
		expected = &jsonapi.Link{}
	)

	if actual, ok := ls.Get(key); actual != nil || ok {
		t.Error("Links.Get should return nil/false when no value is associated with the key")
	}

	ls.Add(key, expected)
	if actual, ok := ls.Get(key); !ok {
		t.Error("expected Links.Get to return true existence check after Links.Add for the given key")
	} else if actual != expected {
		t.Errorf("unexpected value after Links.Get, expected: %s, actual: %s",
			expected, actual)
	}
}

func TestLinks_Get_WhenWrongType(t *testing.T) {
	var (
		ls  = jsonapi.Links{}
		key = "foo"
	)

	ls.AddString(key, "bar")
	v, ok := ls.Get(key)
	if v != nil || ok {
		t.Error("Links.Get should return nil/false when the value associated with the key is not of type *Link")
	}
}

func TestLinks_Get_FromMap(t *testing.T) {
	var (
		ls       = jsonapi.Links{}
		key      = "foo"
		expected = &jsonapi.Link{
			Href: "http://foo.com",
			Meta: map[string]interface{}{
				key: "bar",
			},
		}
	)

	// manually set map[string]interface{} value on underlying Links map to emulate result of JSON umarshaling
	ls[key] = map[string]interface{}{
		"href": expected.Href,
		"meta": expected.Meta,
	}
	if actual, ok := ls.Get(key); actual == nil || !ok {
		t.Error("expected Links.Get to return *Links/true after calling Links.Get on key with underlying map[string]interface{} value")
	} else if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected Links.Get to return correct *Links value, expected: %v, actual: %v", expected, actual)
	}
}

func TestLinks_Delete_WhenNil(t *testing.T) {
	var ls jsonapi.Links = nil
	defer catchPanic(t, "Links", "Delete")
	ls.Delete("foo")
}

func TestLinks_Delete(t *testing.T) {
	var (
		ls     = jsonapi.Links{}
		strKey = "foo"
		objKey = "bar"
	)

	ls.AddString(strKey, "")
	ls.Delete(strKey)
	if actual, ok := ls.GetString(strKey); len(actual) > 0 || ok {
		t.Error("Links.GetString should return empty string/false after calling Links.Delete for the given key")
	}

	ls.Add(objKey, &jsonapi.Link{})
	ls.Delete(objKey)
	if actual, ok := ls.Get(objKey); actual != nil || ok {
		t.Error("Links.Get should return nil/false after calling Links.Delete for the given key")
	}
}
