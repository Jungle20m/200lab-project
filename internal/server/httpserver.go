package server

import (
	"200lab/common"
	usertransport "200lab/internal/modules/user/transport"
	"github.com/gin-gonic/gin"
)

//type User struct {
//	ID         int        `gorm:"column:id"`
//	Name       string     `gorm:"column:name"`
//	Email      string     `gorm:"column:email"`
//	Phone      string     `gorm:"column:phone"`
//	CreateTime *time.Time `gorm:"column:create_time"`
//	UpdateTime *time.Time `gorm:"column:update_time"`
//}

func NewHttpHandler(appCtx common.IAppContext) *gin.Engine {
	//config := appContext.GetConfig()

	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	//// Api logger
	//gin.DisableConsoleColor()
	//f, _ := os.Create(filepath.Join(config.Log.Folder, config.Log.ApiLogFile))
	//gin.DefaultWriter = io.MultiWriter(f)
	//handler.Use(gin.Logger())

	handler.Use(gin.Recovery())

	//// Swagger
	//handler.GET("/grab-electric/v3/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//// Authentication
	//handler.Use(middleware.AuthMW())

	// Router
	v1 := handler.Group("/v1/")
	{
		user := v1.Group("/users")
		user.POST("/", usertransport.GetAll(appCtx))
		//order.GET("/:order_code/customer", orderHttpTransport.GetOrderOfCustomer(appCtx))
		//order.GET("/customer/:customer_code", orderHttpTransport.GetListOrderOfCustomer(appCtx))
	}

	return handler

}
