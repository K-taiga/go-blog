package handler

import (
	"app/model"
	"app/repository"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ArticleCreateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

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

// echo.ContextでHTTPリクエストを取得する
func ArticleCreate(c echo.Context) error {
	// フォームの内容を受け取る構造体を宣言
	var article model.Article

	// type ArticleCreateOutput structを変数に格納
	var out ArticleCreateOutput

	// c.BindでHTTPリクエストの内容をGOにbind
	// modelのarticleの構造に埋め込む
	if err := c.Bind(&article); err != nil {
		c.Logger().Error(err.Error())

		// リクエストの取得に失敗すれば400エラーをJSONで返す
		// http.StatusBadRequestとArticleCreateOutputを返す
		return c.JSON(http.StatusBadRequest, out)
	}

	// バリデーションチェック
	// 構造体のvalidateのキーでvalidationをかける
	if err := c.Validate(&article); err != nil {
		c.Logger().Error(err.Error())

		out.ValidationErrors = article.ValidationErrors(err)

		// エラーがあればUnprocessable Entityで返す
		return c.JSON(http.StatusUnprocessableEntity, out)
	}

	// 登録処理呼び出し
	res, err := repository.ArticleCreate(&article)
	if err != nil {
		c.Logger().Error(err.Error())

		// 500エラー
		return c.JSON(http.StatusInternalServerError, out)
	}

	// 直前のSQLで入ったレコードのIDを取得
	id, _ := res.LastInsertId()

	// articleにIDをセット
	article.ID = int(id)

	// 構造体に保存した記事の内容を格納
	out.Article = &article

	// ここまでくれば201を返す
	return c.JSON(http.StatusCreated, out)
}
