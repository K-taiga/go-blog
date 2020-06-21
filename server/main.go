package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// templateのパス
const tmplPath = "src/template/"

var e = createMux()

func main() {
	// パスとarticleの処理を紐付ける
	e.GET("/", articleIndex)
	e.GET("/new", articleNew)
	e.GET("/:id", articleShow)
	e.GET("/:id/edit", articleEdit)

	// webサーバーをポート8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}

// 返り値はechoのポインタ(実態)
func createMux() *echo.Echo {
	// echoを生成
	e := echo.New()

	// エラーハンドリングのためのmiddleware
	e.Use(middleware.Recover())
	// 各HTTPリクエストに関するログを出すためのmiddleware
	e.Use(middleware.Logger())
	// httpの応答を圧縮するためのmiddleware
	e.Use(middleware.Gzip())

	// 静的ファイルを利用するためのmiddleware jsとかcssとかを置いてあるディレクトリをインポートできる
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	return e
}

// httpリクエストの情報はecho.Contextの構造体に入ってくる
func articleIndex(c echo.Context) error {
	// httpコードが200ならHello Worldの文字列をレスポンス
	// return c.String(http.StatusOK, "Hello, World!")

	// キーがstring,valueがinterfaceのものをdataに実装
	data := map[string]interface{}{
		"Message": "Article Index",
		"Now":     time.Now(),
	}
	return render(c, "article/index.html", data)
}

func articleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}

	return render(c, "article/new.html", data)
}

func articleShow(c echo.Context) error {
	// パスパラメータの:idをstr->intにする
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Show",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/show.html", data)
}

func articleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/edit.html", data)
}

func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	// pongo2 を利用して、テンプレートファイルとデータから HTML を生成している 返り値はbyteとそのエラー
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	// 定義した htmlBlob() 関数を呼び出し、生成された HTML をバイトデータとして受け取る
	b, err := htmlBlob(file, data)

	// エラーチェック
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// ステータスコード 200 で HTML データをレスポンス
	return c.HTMLBlob(http.StatusOK, b)
}
