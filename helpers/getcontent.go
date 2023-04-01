package helpers

import "github.com/labstack/echo"

func GetContentType(c echo.Context) string {
	return c.Request().Header.Get("Content-Type")
}
