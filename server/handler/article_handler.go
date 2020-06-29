package handler

import (
	"app/model"
	"app/repository"
	"fmt"
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
	// "/articles" のパスでリクエストがあったら "/" にリダイレクト
	if c.Request().URL.Path == "/articles" {
		c.Redirect(http.StatusPermanentRedirect, "/")
	}

	// リポジトリのDB取得のindex処理を呼び出し
	articles, err := repository.ArticleListByCursor(0)

	if err != nil {
		c.Logger().Error(err.Error())

		return c.NoContent(http.StatusInternalServerError)
	}

	// 取得した記事の最後のIDをカーソルとして設定
	var cursor int
	if len(articles) != 0 {
		// 配列は0スタートのため
		cursor = articles[len(articles)-1].ID
	}

	data := map[string]interface{}{
		"Articles": articles,
		"Cursor":   cursor,
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

func ArticleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("articleID"))

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

func ArticleDelete(c echo.Context) error {
	// パスパラメータから記事 ID を取得
	// 文字列型で取得されるので、strconv パッケージを利用して数値型
	id, _ := strconv.Atoi(c.Param("articleID"))

	if err := repository.ArticleDelete(id); err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusInternalServerError, "")
	}

	// 成功時はステータスコード 200 を返却
	return c.JSON(http.StatusOK, fmt.Sprintf("Article %d is deleted.", id))
}

func ArticleList(c echo.Context) error {
	cursor, _ := strconv.Atoi(c.QueryParam("cursor"))

	// リポジトリの処理を呼び出して記事の一覧データを取得
	// 引数にカーソルの値を渡して、ID のどの位置から 10 件取得するかを指定
	articles, err := repository.ArticleListByCursor(cursor)

	if err != nil {
		// サーバーのログにエラー内容を出力します。
		c.Logger().Error(err.Error())

		// HTML ではなく JSON 形式でデータのみを返却するため、
		// c.HTMLBlob() ではなく c.JSON() を呼び出す
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, articles)
}

func ArticleShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("articleID"))

	article, err := repository.ArticleGetByID(id)

	if err != nil {
		c.Logger().Error(err.Error())

		return c.NoContent(http.StatusInternalServerError)
	}

	// テンプレートに渡すデータを map に格納
	data := map[string]interface{}{
		"Article": article,
	}

	// テンプレートファイルとデータを指定して HTML を生成し、クライアントに返却
	return render(c, "article/show.html", data)
}
