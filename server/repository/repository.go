package repository

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func SetDB(d *sqlx.DB) {
	// repository パッケージのグローバル変数にdb接続情報をセット
	db = d
}
