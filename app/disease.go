package app

type DiseaseFormCreate struct {
	DiseaseName string `json:"disease_name" validate:"required"`
	DiseaseDesc string `json:"disease_desc" validate:"required"`
}
