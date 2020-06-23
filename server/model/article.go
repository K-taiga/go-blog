package model

//Articleの構造体にIDとdbのカラムのidを紐付けてメタ情報化
type Article struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
