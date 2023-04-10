package business

import (
	"200lab/common"
	usermodel "200lab/internal/modules/user/model"
)

type IGetUserStorage interface {
	GetAllUsers() ([]usermodel.User, error)
}

type getOrderBusiness struct {
	appCtx  common.IAppContext
	storage IGetUserStorage
}

func NewGetOrderBusiness(appCtx common.IAppContext, storage IGetUserStorage) *getOrderBusiness {
	return &getOrderBusiness{
		appCtx:  appCtx,
		storage: storage,
	}
}

func (biz *getOrderBusiness) GetAllUser() ([]usermodel.User, error) {
	return biz.storage.GetAllUsers()
}
