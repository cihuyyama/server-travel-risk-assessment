package app

type EndemicForm struct {
	DiseaseID uint   `json:"disease_id" validate:"required"`
	Province  string `json:"province" validate:"required"`
	RiskLevel string `json:"risk_level" validate:"required"`
}
