package models

import "time"

type LogHistory struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Province    string    `gorm:"column:province" json:"province"`
	DiseaseName string    `gorm:"column:disease_name" json:"disease_name"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedType string    `gorm:"column:updated_type" json:"updated_type"`
}
