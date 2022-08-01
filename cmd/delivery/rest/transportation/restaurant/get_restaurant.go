package restaurant

import (
	"200lab/cmd/delivery/component"
	restaurantbiz "200lab/cmd/delivery/core/business/restaurant"
	"200lab/cmd/delivery/core/storage"
	"fmt"
	"net/http"
)

func GetRestaurant(appCtx *component.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := appCtx.GetMysqlConnector()

		store := storage.NewStorage()
		store.SetMysqlConnector(db)

		biz := restaurantbiz.NewGetRestaurantBusiness(store)

		restaurant, err := biz.GetRestaurant()
		if err != nil {
			fmt.Printf("error: %+v", err)
		}

		fmt.Println(restaurant)
	}
}
