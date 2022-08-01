package restaurant

import (
	model "200lab/cmd/delivery/core/model/restaurant"
	"context"
)

type GetRestaurantStore interface {
	GetBrandByID(ctx context.Context, id int) (*model.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBusiness(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant() (*model.Restaurant, error) {
	restaurant, err := biz.store.GetBrandByID(context.Background(), 1)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}
