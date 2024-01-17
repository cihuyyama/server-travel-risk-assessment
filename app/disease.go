package app

type DiseaseFormCreate struct {
	DiseaseName string `json:"disease_name" validate:"required"`
	DiseaseDesc string `json:"disease_desc" validate:"required"`
}

type TreatmentForm struct {
	DiseaseID   uint   `json:"disease_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type PreventionForm struct {
	DiseaseID   uint   `json:"disease_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
