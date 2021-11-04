package schedule

import (
	"schapi/database"
	"schapi/domain/schedule"
	"schapi/domain/user"
)

type ScheduleRequest struct {
	ID         int    `gorm:"primary_key;not null"       json:"id"`
	Task       string `gorm:"type:varchar(400)"          json:"task"`
	Start_Date string `gorm:"type:datetime"          json:"start_date"`
	End_Date   string `gorm:"type:datetime"          json:"end_date"`
}

// FetchAllschedules は 全てのスケジュール情報を取得する
func FetchAllschedulesRequest(u *user.User) []schedule.Schedule {

	repo := schedule.NewRepository(database.Conn)

	userId := u.Id
	res := repo.FindAllSchedules(userId)

	return res
}

// Addschedule は スケジュールをDBへ登録する
func AddScheduleRequest(r ScheduleRequest, u *user.User) {

	var inputSchedule = schedule.Schedule{
		Task:       r.Task,
		UserId:     u.Id,
		Start_Date: r.Start_Date,
		End_Date:   r.End_Date,
	}

	repo := schedule.NewRepository(database.Conn)
	repo.InsertSchedule(&inputSchedule)

}

// スケジュールを変更する
func ChangeScheduleRequest(r ScheduleRequest) {
	var inputSchedule = schedule.Schedule{
		ID:         r.ID,
		Task:       r.Task,
		Start_Date: r.Start_Date,
		End_Date:   r.End_Date,
	}

	repo := schedule.NewRepository(database.Conn)
	repo.UpdateStateschedule(&inputSchedule)
}

// スケジュールをDBから削除する
func DeleteScheduleRequest(id int) {
	repo := schedule.NewRepository(database.Conn)
	repo.DeleteSchedule(id)

}
