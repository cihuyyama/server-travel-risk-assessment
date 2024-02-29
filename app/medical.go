package app

type MedicalHistoryForm struct {
	UserID               uint   `gorm:"column:user_id" json:"user_id"`
	Age                  string `gorm:"column:age" json:"age"`
	PreexistingCondition string `gorm:"column:preexisting_condition" json:"preexisting_condition"`
	CurrentMedication    string `gorm:"column:current_medication" json:"current_medication"`
	Allergies            string `gorm:"column:allergies" json:"allergies"`
	PreviousVaccination  string `gorm:"column:previous_vaccination" json:"previous_vaccination"`
	Pregnant             string `gorm:"column:pregnant" json:"pregnant"`
	VaccineBcg           bool   `gorm:"column:vaccine_bcg;default:false" json:"vaccine_bcg"`
	VaccineHepatitis     bool   `gorm:"column:vaccine_hepatitis;default:false" json:"vaccine_hepatitis"`
	VaccineDengue        bool   `gorm:"column:vaccine_dengue;default:false" json:"vaccine_dengue"`
}

type MedicalScoreRiskForm struct {
	UserID               uint `gorm:"column:user_id" json:"user_id"`
	PreexistingCondition int  `gorm:"column:preexisting_condition" json:"preexisting_condition"`
	CurrentMedication    int  `gorm:"column:current_medication" json:"current_medication"`
	Allergies            int  `gorm:"column:allergies" json:"allergies"`
	PreviousVaccination  int  `gorm:"column:previous_vaccination" json:"previous_vaccination"`
	Pregnant             int  `gorm:"column:pregnant" json:"pregnant"`
	Age                  int  `gorm:"column:age" json:"age"`
	TotalScore           int  `gorm:"column:total_score" json:"total_score"`
}
