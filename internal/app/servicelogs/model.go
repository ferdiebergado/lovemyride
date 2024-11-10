package servicelogs

import "github.com/ferdiebergado/lovemyride/internal/pkg/db"

type ServiceLog struct {
	db.Model
	Date        string  `json:"date"`
	Mileage     int     `json:"mileage"`
	Description string  `json:"description"`
	VenueID     string  `json:"venue_id"`
	LaborCost   float32 `json:"labor_cost"`
}
