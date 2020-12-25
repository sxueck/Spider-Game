package main

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.GET("/spider/level-1", easy)
}

func easy(c echo.Context) error {
	userAgent := c.Request().UserAgent()
	if !strings.Contains(userAgent, "Mozilla") {
		return c.String(403, "crawlers will be blocked")
	}

	return c.String(200,"congratulations, the plain text is : wokking")
}
