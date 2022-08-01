package restaurant

type Restaurant struct {
	ID            int    `json:"id" db:"id"`
	WorkingSiteID int    `json:"working_site_id" db:"working_site_id"`
	Code          string `json:"code" db:"code"`
	Name          string `json:"name" db:"name"`
}
