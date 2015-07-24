<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/static/logo.png" alt="Volatile Response" title="Volatile Response"><br><br></p>

Volatile Static is a handler for the [Core](https://github.com/volatile/core).  
It serves files from the `static` directory on the `/static` URL path.

## Installation

```Shell
$ go get -u github.com/volatile/static
```

## Usage

```Go
package main

import (
	"github.com/volatile/core"
	"github.com/volatile/static"
)

func main() {
	static.Use(static.DefaultMaxAge)

	core.Run()
}
```

[![GoDoc](https://godoc.org/github.com/volatile/static?status.svg)](https://godoc.org/github.com/volatile/static)

### Cache "max-age"

The `static.Use` method receives a single parameter: the `max-age` (in seconds) for all the static files.
