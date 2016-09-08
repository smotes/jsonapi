package jsonapi_test

import (
	"encoding/json"
	"testing"

	"github.com/smotes/jsonapi"
)

func TestDocument_MarshalJSON(t *testing.T) {
	b, err := json.Marshal(&testDocument)
	if err != nil {
		t.Errorf("unexpected error when marshaling document: %+v", err)
	}

	expected := testDocumentJSON
	actual := string(b)
	if err := compareJSON(expected, actual); err != nil {
		t.Errorf("unexpected error when comparing document JSON: %v", err)
	}
}

func TestDocument_UnmarshalJSON(t *testing.T) {
	doc := jsonapi.Document{}
	if err := json.Unmarshal([]byte(testDocumentJSON), &doc); err != nil {
		t.Errorf("unexpected error when unmarshaling document: %+v", err)
	}
}
