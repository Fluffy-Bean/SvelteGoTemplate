package front

import (
	"embed"

	"github.com/labstack/echo/v4"
)

//go:generate npm run build
//go:embed dist
var dist embed.FS

var DistDir = echo.MustSubFS(dist, "dist")
