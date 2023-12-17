package models

import "time"

type Disease struct {
	ID             uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DiseaseName    string    `json:"disease_name"`
	DiseaseDesc    string    `json:"disease_desc"`
	DiseaseSymptom []Symptom `gorm:"many2many:disease_symptoms;"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
