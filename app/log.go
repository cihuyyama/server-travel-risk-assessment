package app

type LogResult struct {
	ID          uint   `json:"id"`
	Province    string `json:"province"`
	DiseaseName string `json:"disease_name"`
	CreatedDate string `json:"created_date"`
	CreatedTime string `json:"created_time"`
	UpdatedType string `json:"updated_type"`
}
