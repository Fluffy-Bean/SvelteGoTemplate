package front

import (
	"embed"

	"github.com/labstack/echo/v4"
)

//go:embed dist
var dist embed.FS

var DistDir = echo.MustSubFS(dist, "dist")
