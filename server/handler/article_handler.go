package handler

import (
	"app/repository"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func ArticleIndex(c echo.Context) error {
	// repositoryのパッケージからArticleListを取得
	articles, err := repository.ArticleList()
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// キーがstring,valueがinterfaceのものをdataに実装
	data := map[string]interface{}{
		"Message":  "Article Index",
		"Now":      time.Now(),
		"Articles": articles,
	}
	return render(c, "article/index.html", data)
}

func ArticleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}

	return render(c, "article/new.html", data)
}

func ArticleShow(c echo.Context) error {
	// パスパラメータの:idをstr->intにする
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Show",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/show.html", data)
}

func ArticleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/edit.html", data)
}
