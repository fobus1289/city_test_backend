package company

import (
	"github.com/fobus1289/marshrudka/router"
	"v2/app/http/v2/handler"
	"v2/app/http/v2/middleware"
	"v2/app/http/v2/request/company"
)

func User(drive *router.Drive) {

	userRequest := &company.User{}
	userHandler := &handler.User{}
	drive.Dep(userRequest)
	drive.Dep(userHandler)

	userGroup := drive.Group("v2/company/:companyId{n}", middleware.Company)
	{
		userGroup.GET("user", userRequest.All, userHandler.All)
		userGroup.GET("user/:userId{n}", userRequest.FindById, userHandler.FindById)
		userGroup.GET("user-name/:username{s}", userRequest.FindByUsername, userHandler.FindByUsername)
		userGroup.GET("user-fio/:fio{s}", userRequest.FindByFIO, userHandler.FindByFIO)
		userGroup.GET("user/branch/:branchId{n}", userRequest.FindByBranchId, userHandler.FindByBranchId)
		userGroup.POST("user", userRequest.Create, userHandler.Create)
		userGroup.PUT("user/:userId{n}", userRequest.Update, userHandler.Update)
		userGroup.DELETE("user/:userId{n}", userRequest.Delete, userHandler.Delete)
		userGroup.DELETE("user/:userId{n}/role/:roleId{n}", userRequest.DeleteRole, userHandler.DeleteRole)
	}

}
