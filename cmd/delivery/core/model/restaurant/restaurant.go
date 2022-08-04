package restaurant

import (
	"200lab/common"
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	OwnerID         int    `json:"owner_id" db:"owner_id"`
	Name            string `json:"name" db:"name"`
	Address         string `json:"address" db:"address"`
	Logo            string `json:"logo" db:"logo"`
	Cover           string `json:"cover" db:"cover"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantInsert struct {
	common.SQLModel `json:",inline"`
	OwnerID         int    `json:"owner_id" db:"owner_id"`
	Name            string `json:"name" db:"name"`
	Address         string `json:"address" db:"address"`
	Logo            string `json:"logo" db:"logo"`
	Cover           string `json:"cover" db:"cover"`
}

func (RestaurantInsert) TableName() string {
	return Restaurant{}.TableName()
}
