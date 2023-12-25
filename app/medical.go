package app

type MedicalHistoryForm struct {
	UserID               uint   `gorm:"column:user_id" json:"user_id"`
	Age                  string `gorm:"column:age" json:"age"`
	PreexistingCondition string `gorm:"column:preexisting_condition" json:"preexisting_condition"`
	CurrentMedication    string `gorm:"column:current_medication" json:"current_medication"`
	Allergies            string `gorm:"column:allergies" json:"allergies"`
	PreviousVaccination  string `gorm:"column:previous_vaccination" json:"previous_vaccination"`
	Pregnant             string `gorm:"column:pregnant" json:"pregnant"`
	VaccineBcg           bool   `gorm:"column:vaccine_bcg" json:"vaccine_bcg"`
	VaccineHepatitis     bool   `gorm:"column:vaccine_hepatitis" json:"vaccine_hepatitis"`
	VaccineDengue        bool   `gorm:"column:vaccine_dengue" json:"vaccine_dengue"`
}
