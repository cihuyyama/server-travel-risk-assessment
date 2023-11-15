package app

type SymptomFormCreate struct {
	SymptomName string `json:"symptom_name" valid:"required"`
	SymptomChar string `json:"symptom_char" valid:"optional"`
}

type SymptomFormUpdate struct {
	SymptomName string `json:"symptom_name" valid:"required"`
	SymptomChar string `json:"symptom_char" valid:"optional"`
}

type SymptomResult struct {
	ID          string `json:"id"`
	SymptomName string `json:"symptom_name"`
	SymptomChar string `json:"symptom_char"`
}
