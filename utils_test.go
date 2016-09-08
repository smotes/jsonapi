package jsonapi_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func catchPanic(t *testing.T, typ, method string) {
	if r := recover(); r != nil {
		t.Errorf("%s.%s() should not panic when %s is nil", typ, method, typ)
	}
}

func compareJSON(a, b string) error {
	var ar, br interface{}

	err := json.Unmarshal([]byte(a), &ar)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(b), &br)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(ar, br) {
		a = stripJSON(a)
		b = stripJSON(b)
		return fmt.Errorf("\n%s\n!=\n%s", a, b)
	}
	return nil
}

func stripJSON(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	return s
}
