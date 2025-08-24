package example

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goflash/flash/v2"
)

func TestExampleMiddleware(t *testing.T) {
	app := flash.New()

	called := 0

	// Register middleware under test
	app.Use(ExampleMiddleware)

	// Register a simple handler to verify next() is called
	app.GET("/", func(c flash.Ctx) error {
		called++
		return c.String(http.StatusOK, "ok")
	})

	// Perform request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	// Assertions
	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", rec.Code)
	}
	body, _ := io.ReadAll(rec.Body)
	if string(body) != "ok" {
		t.Fatalf("unexpected body: %q", string(body))
	}
	if called != 1 {
		t.Fatalf("handler should be called once, got %d", called)
	}
}
