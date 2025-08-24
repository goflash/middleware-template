module github.com/goflash/example-middleware

go 1.23.0

toolchain go1.23.2

// Uncomment when need to use local GoFlash during development
// replace github.com/goflash/flash/v2 => ../flash

require github.com/goflash/flash/v2 v2.0.0-beta.6

require (
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
)
