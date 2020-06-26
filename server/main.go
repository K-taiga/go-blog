// パッケージ名がmainのパッケージは扱いが特殊で、main関数を定義することでエントリポイント（プログラム実行時の処理開始位置）として使用される
package main

import (
	"app/handler"
	"app/repository"
	"log"
	"os"

	// db
	// blank(_) import => importしたパッケージと依存関係のあるパッケージをimportしてそれを初期化するために必要、依存関係を解決するためのimport
	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"          //sqlを書くデータマッパー型のORM

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

var db *sqlx.DB

func main() {
	db = connectDB()
	repository.SetDB(db)
	e.GET("/", handler.ArticleIndex)
	e.GET("/new", handler.ArticleNew)
	e.GET("/:id", handler.ArticleShow)
	e.GET("/:id/edit", handler.ArticleEdit)
	e.POST("/", handler.ArticleCreate)

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
	// CSRF対策のmiddleware
	e.Use(middleware.CSRF())

	// 静的ファイルを利用するためのmiddleware jsとかcssとかを置いてあるディレクトリをインポートできる
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	return e
}

func connectDB() *sqlx.DB {
	// osパッケージで環境変数を取得
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	// dbにpingが通るかチェック
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}
