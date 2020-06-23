package model

//Articleの構造体にIDとdbのカラムのidをメタ情報化し、sqlxがSQLの実行結果と構造体を紐づける
type Article struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
