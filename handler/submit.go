package handler

import (
	"net/http"
	"strings"
	"url-shortener/storage"
	"url-shortener/util"

	"github.com/labstack/echo/v4"
)

func SubmitHandler(c echo.Context) error {
	url := c.FormValue("url")
	if url == "" {
		return c.String(http.StatusBadRequest, "url is required")
	}

	// 自动补全协议
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	// 生成短链 ID
	id := util.GenerateRandomString(8)

	// 保存
	storage.LinkMap[id] = &storage.Link{
		Id:  id,
		Url: url,
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
