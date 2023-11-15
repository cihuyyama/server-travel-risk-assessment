package models

type Symptom struct {
	ID          uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SymptomName string `json:"symptom_name"`
	SymptomChar string `json:"symptom_char"`
}
