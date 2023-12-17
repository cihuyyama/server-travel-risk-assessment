package models

type MedicalHistory struct {
	ID uint `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}
