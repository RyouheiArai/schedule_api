package schedule

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// FindAllschedules は スケジュールテーブルのレコードを全件取得する
func (r Repository) FindAllSchedules(userId uint64) []Schedule {
	schedules := []Schedule{}

	// select
	r.db.Where("user_id = ?", userId).Order("Start_Date asc").Find(&schedules)

	return schedules
}

// FindSchedule は スケジュールテーブルのレコードを１件取得する
func (r Repository) FindSchedule(scheduleID int) []Schedule {
	schedule := []Schedule{}

	// select
	r.db.First(&schedule, scheduleID)

	return schedule
}

// Insertschedule は スケジュールテーブルにレコードを追加する
func (r Repository) InsertSchedule(schedule *Schedule) {

	// insert
	r.db.Create(&schedule)
}

// UpdateStateschedule は スケジュールテーブルの指定したレコードの状態を変更する
func (r Repository) UpdateStateschedule(schedule *Schedule) {

	var scheduleIDStr = strconv.Itoa(schedule.ID)
	// update
	r.db.Model(&schedule).Where("ID = ?", scheduleIDStr).Update(schedule)

}

// DeleteSchedule は スケジュールテーブルの指定したレコードを削除する
func (r Repository) DeleteSchedule(scheduleID int) {
	schedule := []Schedule{}

	r.db.Delete(&schedule, scheduleID)

}
