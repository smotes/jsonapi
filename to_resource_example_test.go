package jsonapi_test

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/smotes/jsonapi"
)

// satisfy identityReadAdapter on *Person
func (p *Person) GetID() (string, error) {
	return strconv.Itoa(p.ID), nil
}

// satisfy identityReadAdapter on *Person
func (p *Person) GetType() (string, error) {
	return "people", nil
}

// satisfy attributesReadAdapter on *Person
func (p *Person) GetAttributes() (map[string]interface{}, error) {
	as := make(map[string]interface{}, 0)
	as["name"] = p.Name
	as["age"] = p.Age
	return as, nil
}

// satisfy identityReadAdapter on *Article
func (a *Article) GetID() (string, error) {
	return strconv.Itoa(a.ID), nil
}

// satisfy identityReadAdapter on *Article
func (a *Article) GetType() (string, error) {
	return "articles", nil
}

// satisfy attributesReadAdapter on *Article
func (a *Article) GetAttributes() (map[string]interface{}, error) {
	as := make(map[string]interface{}, 0)
	as["title"] = a.Title
	as["body"] = a.Body
	return as, nil
}

// satisfy relationshipsReadAdapter on *Article
func (a *Article) GetRelationships() (jsonapi.Relationships, error) {
	rs := jsonapi.Relationships{}

	// setup author relationship
	r, err := jsonapi.ToResource(a.Author, false)
	if err != nil {
		return nil, err
	}
	d, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	l := jsonapi.Links{}
	l.AddString("self", "http://example.com/articles/1/relationships/author")
	l.AddString("related", "http://example.com/articles/1/author")
	rs.Add("author", &jsonapi.Relationship{
		Data:  d,
		Links: l,
	})

	return rs, nil
}

// satisfy linksReadAdapter on *Article
func (a *Article) GetLinks() (jsonapi.Links, error) {
	ls := jsonapi.Links{}
	ls.AddString("self", "http://example.com/articles/1")
	return ls, nil
}

// satisfy metaReadAdapter on *Article
func (a *Article) GetMeta() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["total"] = 42
	m["type"] = "foo"
	return m, nil
}

// This example includes an integration test with jsonapi.ToResource function
// and the json.Marshal function from the "encoding/json" package.
func ExampleToResource() {
	article := Article{
		ID:    4,
		Title: "some title",
		Body:  "some body",
		Author: &Person{
			ID:   1,
			Name: "John",
			Age:  42,
		},
	}

	// convert article to resource object
	resource, err := jsonapi.ToResource(&article, true)
	if err != nil {
		panic(err)
	}

	// then marshal resource using "encoding/json" package
	b, err := json.Marshal(&resource)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	// Output: {"id":"4","type":"articles","attributes":{"body":"some body","title":"some title"},"relationships":{"author":{"links":{"related":"http://example.com/articles/1/author","self":"http://example.com/articles/1/relationships/author"},"data":{"id":"1","type":"people"}}},"links":{"self":"http://example.com/articles/1"},"meta":{"total":42,"type":"foo"}}
}
