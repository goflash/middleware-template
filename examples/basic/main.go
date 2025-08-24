package main

import (
	"log"
	"net/http"

	"github.com/goflash/example-middleware"
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
