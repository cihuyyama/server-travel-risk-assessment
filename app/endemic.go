package app

type EndemicForm struct {
	DiseaseID uint   `json:"disease_id" validate:"required"`
	Province  string `json:"province" validate:"required"`
	RiskLevel string `json:"risk_level" validate:"required"`
	RiskScore int    `json:"risk_score" validate:"required"`
}

type AppendDiseaseForm struct {
	EndemicID uint `json:"endemic_id" validate:"required"`
	DiseaseID uint `json:"disease_id" validate:"required"`
}
