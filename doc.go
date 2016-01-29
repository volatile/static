/*
Package static is a handler for the core (https://godoc.org/github.com/volatile/core).
It serves files from the "static" directory, wth the "/static" URL path prefix.

Usage

Use adds the handler to the default handlers stack:

	static.Use(static.DefaultMaxAge)
*/
package static
