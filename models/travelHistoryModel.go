package models

import "time"

type TravelHistory struct {
	ID            uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID        uint      `gorm:"column:user_id" json:"user_id"`
	City          string    `gorm:"column:city" json:"city"`
	Province      string    `gorm:"column:province" json:"province"`
	Duration      string    `gorm:"column:duration" json:"duration"`
	TravelPurpose string    `gorm:"column:travel_purpose" json:"travel_purpose"`
	DeparturedAt  string    `gorm:"column:departured_at" json:"departured_at"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type TravelScoreRisk struct {
	ID               uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID           uint   `gorm:"column:user_id" json:"user_id"`
	Duration         int    `gorm:"column:duration" json:"duration"`
	TravelPurpose    int    `gorm:"column:travel_purpose" json:"travel_purpose"`
	DestinationScore int    `gorm:"column:destination_score" json:"destination_score"`
	TotalScore       int    `gorm:"column:total_score" json:"total_score"`
	Categories       string `gorm:"column:categories" json:"categories"`
}
