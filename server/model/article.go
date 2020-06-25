package model

import "time"

//Articleの構造体にIDとdbのカラムのidをメタ情報化し、sqlxがSQLの実行結果と構造体を紐づける
type Article struct {
	ID      int       `db:"id"`
	Title   string    `db:"title"`
	Body    string    `db:"body" form:"body"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}
