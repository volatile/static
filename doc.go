/*
Package static is a handler for the Core.
It serves files from the "static" directory on the "/static" URL path.

Usage

Example:

	package main

	import (
		"github.com/volatile/core"
		"github.com/volatile/static"
	)

	func main() {
		static.Use(static.DefaultMaxAge)

		core.Run()
	}

Cache max-age

The "static.Use" method receives a single parameter: the "max-age" (in seconds) for all resources.
This setting is only used in production environment.
*/
package static
