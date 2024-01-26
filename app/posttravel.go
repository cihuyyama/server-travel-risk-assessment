package app

type PostTravelForm struct {
	Symptoms []uint `json:"symptoms" validate:"required"`
}

type PostTravelResponse struct {
	DiseaseId   uint   `json:"disease_id"`
	DiseaseName string `json:"disease_name"`
	Percentage  int    `json:"percentage"`
}
