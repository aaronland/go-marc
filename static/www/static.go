// package www provides an embedded filesystem for web application static assets (JavaScript, CSS).
package www

import (
	"embed"
)

//go:embed index.html css/*.css javascript/*.js
var FS embed.FS
