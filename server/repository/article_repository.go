package repository

import (
	"app/model"
	"database/sql"
	"math"
	"time"
)

func ArticleListByCursor(cursor int) ([]*model.Article, error) {
	// cursorの値が0以下ならMaxIntにする
	if cursor <= 0 {
		cursor = math.MaxInt32
	}
	// ID の降順に記事データを 10 件取得
	query := `SELECT * FROM articles WHERE id < ? ORDER BY id desc	LIMIT 10`

	// クエリ結果を入れるスライス
	// 10件までのためcapacityを指定
	articles := make([]*model.Article, 0, 10)

	// クエリの実行
	if err := db.Select(&articles, query, cursor); err != nil {
		return nil, err
	}

	return articles, nil

}

func ArticleCreate(article *model.Article) (sql.Result, error) {

	now := time.Now()

	//  *model.Articleでポインタ参照しているのでarticle構造体のフィールドに直接入れられる
	article.Created = now
	article.Updated = now

	query := `INSERT INTO articles (title, body, created, updated)
	VALUES (:title, :body, :created, :updated);`

	// db接続開始
	// tx = 送信機？
	tx := db.MustBegin()

	// クエリ文字列と構造体を引数に渡して SQL を実行します。
	// クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換
	// 構造体タグで指定してあるフィールドが対象（`db:"title"` など）になる
	res, err := tx.NamedExec(query, article)
	if err != nil {
		// 登録時エラーがあればrollback
		tx.Rollback()

		// resにnil　errにerrを返す
		return nil, err
	}

	tx.Commit()

	return res, nil

}
