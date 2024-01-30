package app

type PreTravelProps struct {
}

type DrugConflict struct {
	DiseaseID uint   `json:"disease_id" validate:"required"`
	DrugName  string `json:"drug_name" validate:"required"`
	DrugDesc  string `json:"drug_desc" validate:"required"`
}
