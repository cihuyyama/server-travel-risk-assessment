package app

type TravelHistoryForm struct {
	ID            uint   `json:"id"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Duration      string `json:"duration"`
	TravelPurpose string `json:"travel_purpose"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
