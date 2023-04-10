package transport

import (
	"200lab/common"
	"200lab/internal/modules/user/business"
	"200lab/internal/modules/user/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetByID(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	}
}

func GetAll(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		st := storage.NewMysqlStorage(appCtx.GetDB())
		biz := business.NewGetOrderBusiness(appCtx, st)

		users, err := biz.GetAllUser()
		if err != nil {
			c.JSON(http.StatusOK, "error")
			return
		}

		fmt.Printf("users: %+v\n", users)

		c.JSON(http.StatusOK, "success")
	}
}
