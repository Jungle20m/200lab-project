package restaurant

import (
	"200lab/cmd/delivery/component"
	restaurantbiz "200lab/cmd/delivery/core/business/restaurant"
	restaurantmodel "200lab/cmd/delivery/core/model/restaurant"
	"200lab/cmd/delivery/core/storage"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func InsertRestaurant(appCtx *component.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := restaurantmodel.RestaurantInsert{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			fmt.Println(err)
			return
		}

		db := appCtx.GetMysqlConnector()

		store := storage.NewStorage()
		store.SetMysqlConnector(db)

		biz := restaurantbiz.NewInsertRestaurantBusiness(store)

		err = biz.InsertRestaurant(context.Background(), &body)
		if err != nil {
			fmt.Printf("error: %+v", err)
		}

		fmt.Println("done")
	}
}
