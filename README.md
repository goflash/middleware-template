# Middleware template for the GoFlash web framework

<h1 align="center">
    <a href="https://pkg.go.dev/github.com/goflash/middleware-template/v2@v2.0.0">
        <img src="https://pkg.go.dev/badge/github.com/goflash/middleware-template.svg" alt="Go Reference">
    </a>
    <a href="https://goreportcard.com/report/github.com/goflash/middleware-template">
        <img src="https://img.shields.io/badge/%F0%9F%93%9D%20Go%20Report-A%2B-75C46B?style=flat-square" alt="Go Report Card">
    </a>
    <a href="https://codecov.io/gh/goflash/middleware-template">
        <img src="https://codecov.io/gh/goflash/middleware-template/graph/badge.svg?token=VRHM48HJ5L" alt="Coverage">
    </a>
    <a href="https://github.com/goflash/middleware-template/actions?query=workflow%3ATest">
        <img src="https://img.shields.io/github/actions/workflow/status/goflash/middleware-template/test-coverage.yml?branch=main&label=%F0%9F%A7%AA%20Tests&style=flat-square&color=75C46B" alt="Tests">
    </a>
    <img src="https://img.shields.io/badge/go-1.23%2B-00ADD8?logo=golang" alt="Go Version">
    <a href="https://docs.goflash.dev">
        <img src="https://img.shields.io/badge/%F0%9F%92%A1%20GoFlash-docs-00ACD7.svg?style=flat-square" alt="GoFlash Docs">
    </a>
    <img src="https://img.shields.io/badge/status-stable-green" alt="Status">
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="License">
    <br>
    <div style="text-align:center">
      <a href="https://discord.gg/QHhGHtjjQG">
        <img src="https://dcbadge.limes.pink/api/server/https://discord.gg/QHhGHtjjQG" alt="Discord">
      </a>
    </div>
</h1>

## Features

- Minimal example middleware (`ExampleMiddleware`) showing the pattern
- Runnable example app under `examples/basic`
- Go module with Go 1.23 toolchain
- Ready to copy/fork for new middleware in the GoFlash ecosystem

## Installation

```sh
go get github.com/goflash/middleware-template
```

Go version: requires Go 1.23+. The module sets `go 1.23` and can be used with newer Go versions.

## Quick start

This is the minimal usage from `examples/basic`:

```go
package main

import (
    "log"
    "net/http"

    "github.com/goflash/middleware-template"
    "github.com/goflash/flash/v2"
)

func main() {
    app := flash.New()

    // Use the example middleware
    app.Use(example.ExampleMiddleware)

    // Define a simple route
    app.GET("/", func(c flash.Ctx) error {
        return c.String(http.StatusOK, "ok")
    })

    log.Fatal(http.ListenAndServe(":8080", app))
}
```

## Configuration

This template doesn’t prescribe configuration. Customize the middleware to your needs by editing `ExampleMiddleware` (or replacing it with your own). A typical middleware should:

- Accept and return a `flash.Handler`
- Optionally read/write values on the request context
- Call `next(c)` to pass control to the next handler

## Examples

- `examples/basic`: minimal setup showing how to plug the middleware into a GoFlash app and serve a simple route.

Run the example locally:

```sh
cd examples/basic
go run .
```

Then in another shell:

```sh
curl -s localhost:8080/
```

## Testing

Run tests with coverage:

```sh
./scripts/coverage.sh
```

## Versioning and compatibility

- Module path: `github.com/goflash/middleware-template`
- Requires Go 1.23+

## Contributing

Issues and PRs are welcome. Please run tests before submitting.

## License

MIT
