package models

import "time"

type Disease struct {
	ID             uint         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DiseaseName    string       `json:"disease_name"`
	DiseaseDesc    string       `json:"disease_desc"`
	DiseaseSymptom []Symptom    `gorm:"many2many:disease_symptoms;"`
	Treatment      []Treatment  `gorm:"foreignKey:DiseaseID"`
	Prevention     []Prevention `gorm:"foreignKey:DiseaseID"`
	CreatedAt      time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}

type Treatment struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DiseaseID   uint      `gorm:"column:disease_id" json:"disease_id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Prevention struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DiseaseID   uint      `gorm:"column:disease_id" json:"disease_id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
