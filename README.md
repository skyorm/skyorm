# skyorm

[![GoDoc](https://godoc.org/github.com/skyorm/skyorm?status.svg)](http://godoc.org/github.com/skyorm/skyorm)
[![Go Report Card](https://goreportcard.com/badge/github.com/skyorm/skyorm)](https://goreportcard.com/report/github.com/skyorm/skyorm)
[![go](https://github.com/skyorm/skyorm/actions/workflows/go_test.yml/badge.svg)](https://github.com/skyorm/skyorm/actions/workflows/go_test.yml)

Extremely powerful and developer friendly ORM for go.

## Features

- 🔥 **Code generation** - Skyorm only works with code, generated by [skygen](https://github.com/skyorm/skygen). No reflexion is used.
- 🔥 **Powerful condition builder** - The Cond interface allows you to build comples queries with go code. The conditions are translated to queries in the provider.
- 🔥 **Easy to implement additional providers** - The Provider interface is extremely simple. Implementing a new Provider is like a walk in the park.
- 🔥 **Simple API** - The DB interface has simple, yet powerful, methods to cover database client requirements in modern development.

### Example
```go
package main

import (
	`log`
	"os"

	"github.com/skyorm/examples/models"
	`github.com/skyorm/postgres`
	`github.com/skyorm/skyorm`
)

func main() {
	p, err := postgres.New(os.Getenv("DSN"), log.Default())
	if err != nil {
		log.Fatal(err)
	}
	db := skyorm.NewDB(p)

	m := &models.M{}
	cond := skyorm.Or(
		skyorm.Eq(m.OrmPkProp(), 2),
		skyorm.Eq(m.UintValueProp(), 4),
	)
	list, err := db.Find(models.MStore, cond, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range list {
		model := v.(*models.M)
		log.Println(*model)
	}
}
```
