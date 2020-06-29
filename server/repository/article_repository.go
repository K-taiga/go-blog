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

func ArticleDelete(id int) error {
	query := "DELETE FROM articles WHERE id = ?"

	tx := db.MustBegin()

	if _, err := tx.Exec(query, id); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func ArticleGetByID(id int) (*model.Article, error) {
	query := `SELECT * FROM articles WHERE id = ?;`

	// 複数件取得の場合はスライスでしたが、一件取得の場合は構造体
	var article model.Article

	// 複数件の取得の場合は db.Select() でしたが、一件取得の場合は db.Get()
	if err := db.Get(&article, query, id); err != nil {
		return nil, err
	}

	return &article, nil
}

func ArticleUpdate(article *model.Article) (sql.Result, error) {
	now := time.Now()

	article.Updated = now

	// クエリ文字列を生成します。
	query := `UPDATE articles SET title = :title, body = :body, updated = :updated WHERE id = :id;`

	tx := db.MustBegin()

	// クエリ文字列と引数で渡ってきた構造体を指定して、SQL を実行
	// クエリ文字列内の :title, :body, :id には、
	// 第 2 引数の Article 構造体の Title, Body, ID が bind されます。
	// 構造体に db タグで指定した値が紐付け
	res, err := tx.NamedExec(query, article)

	if err != nil {
		tx.Rollback()

		return nil, err
	}

	tx.Commit()

	return res, nil
}
