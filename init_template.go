//go:build embedNgingPluginTemplate

package collector

import (
	"embed"
)

//go:embed template
var TemplateFS embed.FS

//go:embed public
var AssetsFS embed.FS
