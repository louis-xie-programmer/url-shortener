package handler

import (
	"net/http"
	"url-shortener/storage"

	"github.com/labstack/echo/v4"
)

func RedirectHandler(c echo.Context) error {
	id := c.Param("id")

	link, ok := storage.LinkMap[id]
	if !ok {
		return c.String(http.StatusNotFound, "short link not found")
	}

	return c.Redirect(http.StatusMovedPermanently, link.Url)
}
