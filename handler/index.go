package handler

import (
	"fmt"
	"net/http"
	"url-shortener/storage"

	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	html := `
		<h1>Go 短链接服务</h1>
		<form method="POST" action="/submit">
			<input name="url" placeholder="请输入一个 URL" style="width: 300px;" />
			<button type="submit">生成短链</button>
		</form>
		<hr>
		<h3>已生成的短链：</h3>
		<ul>
	`

	for _, link := range storage.LinkMap {
		html += fmt.Sprintf("<li><a href='/%s'>/%s</a> → %s</li>",
			link.Id, link.Id, link.Url)
	}

	html += "</ul>"

	return c.HTML(http.StatusOK, html)
}
