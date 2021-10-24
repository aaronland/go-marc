// package html provides an embedded filesystem for web application templates.
package html

import (
	"embed"
)

//go:embed *.html
var FS embed.FS
