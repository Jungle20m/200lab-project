package restaurant

import (
	restaurentmodel "200lab/cmd/delivery/core/model/restaurant"
	"context"
)

type InsertRestaurantStore interface {
	InsertRestaurant(ctx context.Context, data *restaurentmodel.RestaurantInsert) error
}

type insertRestaurantBiz struct {
	store InsertRestaurantStore
}

func NewInsertRestaurantBusiness(store InsertRestaurantStore) *insertRestaurantBiz {
	return &insertRestaurantBiz{store: store}
}

func (biz *insertRestaurantBiz) InsertRestaurant(ctx context.Context, data *restaurentmodel.RestaurantInsert) error {
	err := biz.store.InsertRestaurant(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
