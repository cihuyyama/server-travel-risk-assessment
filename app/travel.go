package app

type TravelHistoryForm struct {
	ID            uint   `json:"id"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Duration      string `json:"duration"`
	TravelPurpose string `json:"travel_purpose"`
	DeparturedAt  string `json:"departured_at"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
