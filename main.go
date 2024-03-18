package main

import (
	"fmt"

	"github.com/Fluffy-Bean/SvelteGoTemplate/front"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Message struct {
	Name string `json:"name"`
}

func main() {
	r := echo.New()

	r.Use(middleware.Logger())
	r.Use(middleware.Recover())
	r.Use(middleware.Gzip())

	r.StaticFS("/", front.DistDir)

	r.POST("/api", func(c echo.Context) error {
		var form Message
		if err := c.Bind(&form); err != nil {
			return err
		}
		return c.JSON(200, map[string]string{
			"message": fmt.Sprintf("Hello, %s!", form.Name),
		})
	})

	r.Logger.Fatal(r.Start(":8080"))
}
