[![Continuous Integration](https://github.com/MajorLettuce/go-locale/actions/workflows/ci.yml/badge.svg)](https://github.com/MajorLettuce/go-locale/actions/workflows/ci.yml)
![Coverage](https://img.shields.io/badge/Coverage-71.8%25-brightgreen)
[![Go Reference](https://pkg.go.dev/badge/github.com/MajorLettuce/go-locale.svg)](https://pkg.go.dev/github.com/MajorLettuce/go-locale)

# go-locale

A [fluent API](https://www.martinfowler.com/bliki/FluentInterface.html) wrapper
for [go-i18n](https://github.com/nicksnyder/go-i18n) module.

## Install

```
go get github.com/MajorLettuce/go-locale
```

## Usage

Here is a complete example on how to use this package:

```go
package main

import (
	"embed"
	"fmt"

	"github.com/MajorLettuce/go-locale"
)

// Embed FS variable must be defined outside of a function
// Make sure 'locale' directory exists in your project directory:
//
// <project root>
// |-- main.go
// |-- locale/
//   |-- en.yml  :  greeting: Hello, {{.name}}!
//   |-- es.yml  :  greeting: Hola, {{.name}}!

//go:embed locale/*.yml
var localeFS embed.FS

func main() {
	// Initialize reads *.yml files from 'locale' directory
	err := locale.Initialize("en", localeFS, "locale")
	if err != nil {
		panic(err)
	}

	// Now you can read locale messages globally
	fmt.Println(
		locale.Message("greeting"), // Hello, <no value>!
	)

	// You can specify message data and language fluently
	fmt.Println(
		locale.Message("greeting").In("es").With("name", "Alex"), // Hola, Alex!
	)
}

```
