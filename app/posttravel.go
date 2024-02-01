package app

type PostTravelForm struct {
	Symptoms []uint `json:"symptoms" validate:"required"`
}

type PostTravelResponse struct {
	DiseaseId   uint   `json:"disease_id"`
	DiseaseName string `json:"disease_name"`
	DiseaseDesc string `json:"disease_desc"`
	Percentage  int    `json:"percentage"`
}
