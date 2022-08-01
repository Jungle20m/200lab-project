package storage

import (
	"context"

	model "200lab/cmd/delivery/core/model/restaurant"
)

func (s *storage) GetBrandByID(ctx context.Context, id int) (*model.Restaurant, error) {
	restaurant := model.Restaurant{}
	if err := s.db.Get(&restaurant, "SELECT id, working_site_id, code, name FROM brands WHERE id=?", id); err != nil {
		return nil, err
	}
	return &restaurant, nil
}
