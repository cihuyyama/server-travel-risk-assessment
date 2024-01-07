package models

import "time"

type Endemicity struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DiseaseID uint      `gorm:"column:disease_id" json:"disease_id"`
	Province  string    `gorm:"column:country_name" json:"country_name"`
	RiskLevel string    `gorm:"column:risk_level" json:"risk_level"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
