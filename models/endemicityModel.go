package models

import "time"

type Endemicity struct {
	ID             uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DiseaseEndemic []Disease `gorm:"many2many:disease_endemic;"`
	Province       string    `gorm:"column:country_name" json:"country_name"`
	RiskLevel      string    `gorm:"column:risk_level" json:"risk_level"`
	RiskScore      int       `gorm:"column:risk_score;default:70" json:"risk_score"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
