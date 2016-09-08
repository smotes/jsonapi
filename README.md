# jsonapi

Package jsonapi converts Go data to and from the [JSON API](http://jsonapi.org) format.

```
go get -u github.com/smotes/jsonapi
```

## Overview

The JSON API specification has strict requirements on the structure of any JSON request/response containing data, a 
structure most likely varying drastically from any backend data structures. This package solely aims to provide 
utilities around converting said backend data to the specification's structure in a type-safe manner, while also
assuming nothing about the JSON encoding/decoding package used.

### Type safe

Implement multiple (mostly optional) interfaces to convert your data to and from JSON API documents and resources.

```go
type Person struct {
	ID   int
	Name string
	Age  int
}

func (p *Person) GetID() (string, error) {
	return strconv.Itoa(p.ID), nil
}

func (p *Person) GetType() (string, error) {
	return "people", nil
}
```

### Use any JSON package

This package converts your data to and from a common Resource struct, which is compatible with any third-party JSON 
package with the "encoding/json" API, or can work with byte slices.

```go
var person = &Person{ID: 1, Name: "John", Age: 42}

resource, err := jsonapi.ToResource(&person, false)
handle(err)

b, err := json.Marshal(&resource)
handle(err)

fmt.Println(string(b))
// {"id": "1", "type": "people"}
```

## Tested examples

Tested, detailed examples are included on the ToResource and FromResource functions in the [godocs](https://godoc.org/github.com/smotes/jsonapi).

## Contributing

* Fork the repository.
* Code your changes.
* If applicable, write tests and documentation for the new functionality (please ensure all tests pass and have 100% coverage).
* Raise a new pull request with a short description.
