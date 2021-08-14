package main

import (
	// ロギングを行うパッケージ
	"log"

	// HTTPを扱うパッケージ

	// Gin
	"github.com/gin-gonic/gin"

	// MySQL用ドライバ
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// コントローラー
	controller "SCH/controllers/controller"
)

func main() {
	// サーバーを起動する
	serve()
}

func serve() {
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェア を保有しています
	router := gin.Default()

	// // 全てのスケジュールのJSONを返す
	router.GET("/fetchAllschedules", controller.FetchAllSchedules)

	// // １つのスケジュールの状態のJSONを返す
	router.GET("/fetchschedule", controller.FindSchedule)

	// // スケジュールをDBへ登録する
	router.POST("/addschedule", controller.AddSchedule)

	// スケジュールを変更する
	router.POST("/changeschedule", controller.ChangeSchedule)

	// // スケジュールを削除する
	router.POST("/deleteschedule", controller.DeleteSchedule)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
