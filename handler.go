package static

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/volatile/core"
)

const (
	// DefaultMaxAge provides a default caching value of 1 hour.
	DefaultMaxAge = 1 * time.Hour

	assetsDir = "static"
)

var (
	fs = http.FileServer(http.Dir(assetsDir))
)

// Use tells the core to use this handler.
func Use(maxAge time.Duration) {
	maxAgeString := fmt.Sprintf("%.f", maxAge.Seconds())
	core.Use(func(c *core.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/"+assetsDir) {
			if core.Production {
				c.ResponseWriter.Header().Set("Cache-Control", "public, max-age="+maxAgeString)
			}
			http.StripPrefix("/"+assetsDir, fs).ServeHTTP(c.ResponseWriter, c.Request)
		} else {
			c.Next()
		}
	})
}
