package database

import (
	// フォーマットI/O

	"log"

	"os"

	// Go言語のORM
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// エンティティ(データベースのテーブルの行に対応)
	"schapi/domain/user"

	// エンティティ(データベースのテーブルの行に対応)
	"schapi/domain/schedule"
)

var Conn *gorm.DB

// DB接続する
func Initialize() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	DBMS := os.Getenv("DB_MS")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := os.Getenv("DB_PROTOCOL")
	DBNAME := os.Getenv("DB_NAME")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Local"
	conn, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		// 処理停止とランタイムエラー出力
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	conn.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	conn.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	conn.AutoMigrate(&schedule.Schedule{}, &user.User{})
	log.Print("db connected: ", &conn)

	Conn = conn
}
