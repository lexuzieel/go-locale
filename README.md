[![Continuous Integration](https://github.com/avixFF/go-locale/actions/workflows/ci.yml/badge.svg)](https://github.com/avixFF/go-locale/actions/workflows/ci.yml)
![Coverage](https://img.shields.io/badge/Coverage-75.4%25-brightgreen)
![Latest version](https://img.shields.io/badge/Latest%20version-v0.2.0-blue)
[![Go Reference](https://pkg.go.dev/badge/github.com/avixFF/go-locale.svg)](https://pkg.go.dev/github.com/avixFF/go-locale)

# go-locale

A [fluent API](https://www.martinfowler.com/bliki/FluentInterface.html) wrapper
for [go-i18n](https://github.com/nicksnyder/go-i18n) module.

## Install

```
go get github.com/avixFF/go-locale
```

## Usage

Here is a complete example on how to use this package:

```go
package main

import (
	"embed"
	"fmt"

	"github.com/avixFF/go-locale"
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
