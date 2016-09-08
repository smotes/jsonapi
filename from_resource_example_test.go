package jsonapi_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/smotes/jsonapi"
)

// satisfy identityWriteAdapter on *Person
func (p *Person) SetID(id string) error {
	temp, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	p.ID = temp
	return nil
}

// satisfy identityWriteAdapter on *Person
func (p *Person) SetType(typ string) error {
	if typ != "people" {
		return errors.New("type should equal people")
	}
	return nil
}

// satisfy identityWriteAdapter on *Article
func (a *Article) SetID(id string) error {
	temp, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	a.ID = temp
	return nil
}

// satisfy identityWriteAdapter on *Article
func (a *Article) SetType(typ string) error {
	if typ != "articles" {
		return errors.New("type should equal articles")
	}
	return nil
}

// satisfy attributesWriteAdapter on *Article
func (a *Article) SetAttributes(as map[string]interface{}) error {
	if v, ok := as["title"]; ok {
		if v, ok := v.(string); ok {
			a.Title = v
		}
	}

	if v, ok := as["body"]; ok {
		if v, ok := v.(string); ok {
			a.Body = v
		}
	}

	if v, ok := as["created"]; ok {
		if v, ok := v.(string); ok {
			temp, err := time.Parse(time.RFC3339Nano, v)
			if err != nil {
				return err
			}
			a.Created = temp
		}
	}

	if v, ok := as["updated"]; ok {
		if v, ok := v.(string); ok {
			temp, err := time.Parse(time.RFC3339Nano, v)
			if err != nil {
				return err
			}
			a.Updated = temp
		}
	}

	return nil
}

// This example includes an integration test with jsonapi.FromResource function and
// the json.Unmarshal function from the "encoding/json" package.
func ExampleFromResource() {
	testArticleJSON = fmt.Sprintf(`
	{
		"id": "1",
		"type": "articles",
		"attributes": {
			"title": "JSON API paints my bikeshed!",
			"body": "The shortest article. Ever.",
			"created": "%s",
			"updated": "%s"
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
	`, now.Format(time.RFC3339Nano), now.Format(time.RFC3339Nano))

	p := Person{}
	article := Article{}
	articleResource := jsonapi.Resource{}
	personResource := jsonapi.Resource{}

	// first, unmarshal and convert article resource to article
	if err := json.Unmarshal([]byte(testArticleJSON), &articleResource); err != nil {
		panic(err)
	}
	if err := jsonapi.FromResource(&article, &articleResource, true); err != nil {
		panic(err)
	}

	// person resource in author relationship must be unmarshaled separately afterwards
	if rel, ok := articleResource.Relationships.Get("author"); ok {
		if err := json.Unmarshal(rel.Data, &personResource); err != nil {
			panic(err)
		}
		if err := jsonapi.FromResource(&p, &personResource, true); err != nil {
			panic(err)
		}
	}
	article.Author = &p

	fmt.Println(article.ID, article.Title, article.Author.ID)
	// Output: 1 JSON API paints my bikeshed! 42
}
