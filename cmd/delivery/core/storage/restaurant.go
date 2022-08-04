package storage

import (
	"context"

	restaurentmodel "200lab/cmd/delivery/core/model/restaurant"
)

func (s *storage) GetRestaurantByID(ctx context.Context, id int) (*restaurentmodel.Restaurant, error) {
	restaurant := restaurentmodel.Restaurant{}
	if err := s.db.Get(&restaurant, "SELECT id, working_site_id, code, name FROM brands WHERE id=?", id); err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func (s *storage) InsertRestaurant(ctx context.Context, data *restaurentmodel.RestaurantInsert) error {
	query := `
				INSERT INTO restaurants(owner_id, name, address, logo, cover, status)
				VALUES (:owner_id, :name, :address, :logo, :cover, :status)
			 `
	s.db.NamedExec(query, ())
	return nil
}
