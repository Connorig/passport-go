package static

import "embed"

//go:embed assets/* device.png favicon.ico index.html
var Static embed.FS
