package restaurant

import (
	model "200lab/cmd/delivery/core/model/restaurant"
	"context"
)

type SelectRestaurantStore interface {
	GetRestaurantByID(ctx context.Context, id int) (*model.Restaurant, error)
}

type selectRestaurantBiz struct {
	store SelectRestaurantStore
}

func NewSelectRestaurantBusiness(store SelectRestaurantStore) *selectRestaurantBiz {
	return &selectRestaurantBiz{store: store}
}

func (biz *selectRestaurantBiz) GetRestaurant(ctx context.Context) (*model.Restaurant, error) {
	restaurant, err := biz.store.GetRestaurantByID(ctx, 1)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}
