// Package jsonapi provides utilities for converting Go data to/from the JSON API format detailed at http://jsonapi.org.
//
// Consider an example, some structs representing a simple backend for a blog engine:
//
//	type Person {
//		ID   int
//		Name string
//		Age  int
//	}
//
// 	type Article struct {
//		ID      int
//		Title   string
//		Body    string
// 		Created time.Time
// 		Updated time.Time
//		Author  *Person
//	}
//
// The article data could be represented in a JSON format in compliance with the JSON API specification like so:
//
//	{
//		"id": "...",
//		"type": "articles",
//		"attributes": {
//			"title": "...",
//			"body": "...",
//			"created": "...",
//			"updated": "..."
// 		},
//		"relationships": {
//			"author": { ... }
// 		},
//		"links": { ... },
//		"meta": { ... }
//	}
//
// In order to expose the article's data in the desired format, the some methods must be implemented to satisfy
// the expected interfaces for converting to/from resources. For a full list of supported interfaces as well as
// tested examples, please see the documentation for ToResource and FromResource.
package jsonapi // import "github.com/smotes/jsonapi"
