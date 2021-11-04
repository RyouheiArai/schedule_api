package scheduleapi

import "github.com/gin-gonic/gin"

func SetupRoute(api *gin.RouterGroup) {

	// // １つのスケジュールの状態のJSONを返す
	api.GET("/fetchAllschedules", FindAllSchedule)

	// // スケジュールをDBへ登録する
	api.POST("/addschedule", AddSchedule)

	// スケジュールを変更する
	api.POST("/changeschedule", ChangeSchedule)

	// // スケジュールを削除する
	api.POST("/deleteschedule", Deleteschedule)
}
