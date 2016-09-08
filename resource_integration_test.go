package jsonapi_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/smotes/jsonapi"
)

func TestResource_MarshalJSON(t *testing.T) {
	r, err := jsonapi.ToResource(&testArticle, true)
	if err != nil {
		t.Errorf("unexpected error when converting article struct to resource: %+v", err)
		return
	}

	b, err := json.Marshal(&r)
	if err != nil {
		t.Errorf("unexpected error when marshaling resource to json: %+v", err)
		return
	}

	actual := string(b)
	expected := testArticleJSON

	if err := compareJSON(expected, actual); err != nil {
		t.Errorf("%+v", err)
	}
}

func TestResource_UnmarshalJSON(t *testing.T) {
	p := Person{}
	r1 := jsonapi.Resource{}
	r2 := jsonapi.Resource{}
	expected := testArticle
	actual := Article{}

	if err := json.Unmarshal([]byte(testArticleJSON), &r1); err != nil {
		t.Errorf("unexpected error when unmarshaling resource from json: %+v", err)
		return
	}
	if err := jsonapi.FromResource(&actual, &r1, true); err != nil {
		t.Errorf("unexpected error when converting resource to article struct: %+v", err)
		return
	}

	// relationship resource(s) must be unmarshaled separately
	if rel, ok := r1.Relationships.Get("author"); ok {
		if err := json.Unmarshal(rel.Data, &r2); err != nil {
			t.Errorf("unexpected error when unmarshaling resource relationship from json: %+v", err)
			return
		}
	}
	if err := jsonapi.FromResource(&p, &r2, true); err != nil {
		t.Errorf("unexpected error when converting resource to person struct: %+v", err)
		return
	}
	actual.Author = &p

	if !reflect.DeepEqual(expected.Author, actual.Author) {
		t.Errorf("article.Author from unmarshaling resource does not match expected value:\n- expected:\t%+v\n- actual\t%+v",
			expected, actual)
	}

	expected.Author = nil
	actual.Author = nil
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("article from unmarshaling resource does not match expected value:\n- expected:\t%+v\n- actual\t%+v",
			expected, actual)
	}
}
