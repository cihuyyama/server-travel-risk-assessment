package app

import "travel-risk-assessment/models"

type PreTravelProps struct {
	Disease          []models.Disease    `json:"disease" validate:"required"`
	Province         string              `json:"province" validate:"required"`
	RiskLevel        string              `json:"risk_level" validate:"required"`
	AllergyRiskLevel string              `json:"allergy_risk_level" validate:"required"`
	DrugConflicts    []DrugConflict      `json:"drug_conflicts" validate:"required"`
	Prevention       []models.Prevention `json:"prevention" validate:"required"`
	Treatment        []models.Treatment  `json:"treatment" validate:"required"`
}

type DrugConflict struct {
	DiseaseID uint   `json:"disease_id" validate:"required"`
	DrugName  string `json:"drug_name" validate:"required"`
	DrugDesc  string `json:"drug_desc" validate:"required"`
}
