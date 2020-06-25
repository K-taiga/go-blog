package repository

import (
	"app/model"
	"database/sql"
	"time"
)

func ArticleList() ([]*model.Article, error) {
	query := `SELECT * FROM articles;`

	// データベースから取得した値を格納する変数を宣言
	// modelのArticle構造体を格納する配列
	var articles []*model.Article

	// Query を実行して、取得した値を変数に格納
	if err := db.Select(&articles, query); err != nil {
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
