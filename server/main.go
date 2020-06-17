package main

import (
	"net/http"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// templateのパス
const tmplPath = "src/template"

var e = createMux()

func main() {
	// `/`のパスとarticleIndexの処理を紐付ける
	e.GET("/", articleIndex)

	// webサーバーをポート8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}

// 返り値はechoのポインタ(実態)
func createMux() *echo.Echo {
	// echoを生成
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

// httpリクエストの情報はecho.Contextの構造体に入ってくる
func articleIndex(c echo.Context) error {
	// httpコードが200ならHello Worldの文字列をレスポンス
	// return c.String(http.StatusOK, "Hello, World!")

	data := map[string]interface{}{
		"Message": "Hello, World!",
		"Now":     time.Now(),
	}
	return render(c, "article/index.html", data)
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
