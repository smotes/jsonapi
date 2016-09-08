package jsonapi_test

import (
	"errors"
	"fmt"

	"github.com/smotes/jsonapi"
)

// Person represents an person in an example blog engine, used in examples and integration tests.
type Person struct {
	ID   int
	Name string
	Age  int
}

// Article represents an article in an example blog engine, used in examples and integration tests.
type Article struct {
	ID     int
	Title  string
	Body   string
	Author *Person
}

// various singleton fixtures used in unit/integration tests, and examples.

// common
var (
	testErr = errors.New("test error")
)

// meta
var (
	testMetaJSON = `{"test": "foo"}`
	testMeta     = map[string]interface{}{
		"test": "foo",
	}
)

// links
var (
	testHref      = "http://test.com"
	testLinksJSON = fmt.Sprintf(`
	{
		"test": "%s",
		"test2": {
			"href": "%s",
			"meta": %s
		}
	}`, testHref, testHref, testMetaJSON)
	testLinks = func() jsonapi.Links {
		ls := jsonapi.Links{}
		ls.AddString("test", testHref)
		ls.Add("test2", &jsonapi.Link{
			Href: testHref,
			Meta: testMeta,
		})
		return ls
	}()
)

// errors
var (
	testErrorJSON = fmt.Sprintf(`
	{
		"id": "test",
		"status": "test",
		"code": "test",
		"title": "test",
		"detail": "test",
		"links": %s,
		"meta": %s
	}`, testLinksJSON, testMetaJSON)
	testError = jsonapi.Error{
		ID:     "test",
		Status: "test",
		Code:   "test",
		Title:  "test",
		Detail: "test",
		Links:  testLinks,
		Meta:   testMeta,
	}
)

// jsonapi info
var (
	testInfoJSON = fmt.Sprintf(`
	{
		"version": "test",
		"meta": %s
	}`, testMetaJSON)
	testInfo = jsonapi.Info{
		Version: "test",
		Meta:    testMeta,
	}
)

// document
var (
	testDocumentJSON = fmt.Sprintf(`
	{
		"errors": [%s],
		"jsonapi": %s,
		"links": %s,
		"meta": %s
	}`, testErrorJSON, testInfoJSON, testLinksJSON, testMetaJSON)
	testDocument = func() jsonapi.Document {
		doc := jsonapi.Document{
			Errors: []jsonapi.Error{testError},
			Info:   &testInfo,
			Links:  testLinks,
			Meta:   testMeta,
		}
		return doc
	}()
)

// article resource
var (
	testPerson = Person{
		ID: 42,
	}
	testArticle = Article{
		ID:     1,
		Title:  "JSON API paints my bikeshed!",
		Body:   "The shortest article. Ever.",
		Author: &testPerson,
	}
	testArticleJSON = `
	{
		"id": "1",
		"type": "articles",
		"attributes": {
			"title": "JSON API paints my bikeshed!",
			"body": "The shortest article. Ever."
		},
		"relationships": {
			"author": {
				"data": {
					"id": "42",
					"type": "people"
				},
				"links": {
					"self": "http://example.com/articles/1/relationships/author",
					"related": "http://example.com/articles/1/author"
				}
			}
		},
		"links": {
			"self": "http://example.com/articles/1"
		},
		"meta": {
			"total": 42,
			"type": "foo"
		}
	}
	`
)
