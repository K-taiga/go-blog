package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

//Articleの構造体にIDとdbのカラムのidをメタ情報化し、sqlxがSQLの実行結果と構造体を紐づける
type Article struct {
	ID      int       `db:"id"`
	Title   string    `db:"title" form:"title" validate:"required,max=50"`
	Body    string    `db:"body" form:"body" validate:"required"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}

// ValidatonErrors
// 返り値がerr型はerror型　fmt.Stringerに似た組み込みのインタフェース　stringの配列
// 返り値が複数あれば(,)で指定する
func (a *Article) ValidationErrors(err error) []string {
	var errMessages []string
	// 複数のエラーが発生する可能性があるためforで回す
	for _, err := range err.(validator.ValidationErrors) {

		var message string

		switch err.Field() {
		case "Title":
			// どのvalidationでエラーになったかキーで判断
			switch err.Tag() {
			case "required":
				message = "タイトルは必須です。"
			case "max":
				message = "タイトルは最大50文字です。"
			}
		case "Body":
			message = "本文は必須です。"
		}

		// メッセージをerrMessageのスライスに追加して代入
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}
	return errMessages
}
