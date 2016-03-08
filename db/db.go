package db

import (
	// initialize SQLite3 Driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
)

// 外部からアクセスするためのプレースホルダー
var shared *genmai.DB

// GetClient 外部へdb clientを提供する
func GetClient() *genmai.DB {
	if shared == nil {
		panic("not connected to DB, use Dial")
	}
	return shared
}

// Dial Sharedにdb clientを格納させ、それ自身を返す
func Dial(dialect genmai.Dialect, dsn string) {
	if shared != nil {
		return
	}

	db, err := genmai.New(dialect, dsn)

	if err != nil {
		// DB初期化はアプリケーション立ち上げ時に行われるので、
		// 失敗した場合はアプリケーションが落ちるべき
		panic(err)
	}

	shared = db
}

// Close db clientを閉じる
func Close() {
	if shared == nil {
		return
	}
	shared.Close()
}
