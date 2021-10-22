package static

import (
	"embed"
)

//go:embed css/*.css javascript/*.js
var FS embed.FS
