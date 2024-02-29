package models

import "time"

type MedicalHistory struct {
	ID                   uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID               uint      `gorm:"column:user_id" json:"user_id"`
	Age                  string    `gorm:"column:age" json:"age"`
	PreexistingCondition string    `gorm:"column:preexisting_condition" json:"preexisting_condition"`
	CurrentMedication    string    `gorm:"column:current_medication" json:"current_medication"`
	Allergies            string    `gorm:"column:allergies" json:"allergies"`
	PreviousVaccination  string    `gorm:"column:previous_vaccination" json:"previous_vaccination"`
	Pregnant             string    `gorm:"column:pregnant" json:"pregnant"`
	VaccineBcg           bool      `gorm:"column:vaccine_bcg" json:"vaccine_bcg"`
	VaccineHepatitis     bool      `gorm:"column:vaccine_hepatitis" json:"vaccine_hepatitis"`
	VaccineDengue        bool      `gorm:"column:vaccine_dengue" json:"vaccine_dengue"`
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MedicalScoreRisk struct {
	ID                   uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID               uint   `gorm:"column:user_id" json:"user_id"`
	PreexistingCondition int    `gorm:"column:preexisting_condition" json:"preexisting_condition"`
	CurrentMedication    int    `gorm:"column:current_medication" json:"current_medication"`
	Allergies            int    `gorm:"column:allergies" json:"allergies"`
	PreviousVaccination  int    `gorm:"column:previous_vaccination" json:"previous_vaccination"`
	Pregnant             int    `gorm:"column:pregnant" json:"pregnant"`
	Age                  int    `gorm:"column:age" json:"age"`
	TotalScore           int    `gorm:"column:total_score" json:"total_score"`
	Categories           string `gorm:"column:categories" json:"categories"`
}
