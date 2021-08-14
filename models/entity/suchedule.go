package entity

// Schedule はテーブルのモデル
type Schedule struct {
	ID         int    `gorm:"primary_key;not null"       json:"id"`
	Company    string `gorm:"type:varchar(200);not null" json:"company"`
	Task       string `gorm:"type:varchar(400)"          json:"task"`
	Start_Date string `gorm:"type:datetime"          json:"start_date"`
	End_Date   string `gorm:"type:datetime"          json:"end_date"`
}
