package transport

import (
	"200lab/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetByID(appCtx common.IAppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetDB()

		fmt.Println(db)

		c.JSON(http.StatusOK, "success")
	}
}
