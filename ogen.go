// Package ogen implements OpenAPI v3 code generation.
package ogen

//go:generate go run ./cmd/ogen --schema _testdata/sample_1.json --target internal/sample_api --clean --debug.noerr
//go:generate go run ./cmd/ogen --schema _testdata/techempower.json --target internal/techempower --package techempower --clean
