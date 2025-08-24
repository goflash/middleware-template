module example-basic

go 1.23.0

toolchain go1.23.2

// Using local middleware codebase in local examples
replace github.com/goflash/example-template => ../..

require (
	github.com/goflash/example-template v0.0.0-00010101000000-000000000000
	github.com/goflash/flash/v2 v2.0.0-beta.6
)

require (
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
)
