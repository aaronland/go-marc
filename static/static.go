// package static provides an embedded filesystem for web application static assets (JavaScript, CSS).
package static

import (
	"embed"
)

//go:embed css/*.css javascript/*.js
var FS embed.FS
