package admin

import (
	"github.com/fobus1289/marshrudka/router"
	"v2/app/http/v2/handler"
	"v2/app/http/v2/middleware"
	"v2/app/http/v2/request/admin"
)

func User(drive *router.Drive) {

	userRequest := &admin.User{}
	userHandler := &handler.User{}

	drive.Dep(userRequest)
	drive.Dep(userHandler)

	userGroup := drive.Group("v2/admin", middleware.Admin)
	{
		userGroup.GET("users", userRequest.All, userHandler.All)
		userGroup.GET("user/:id{n}", userRequest.FindById, userHandler.FindById)
		userGroup.POST("user", userRequest.Create, userHandler.Create)
		userGroup.PUT("user/:id{n}", userRequest.Update, userHandler.Update)
		userGroup.DELETE("user/:id{n}", userRequest.Delete, userHandler.Delete)
		userGroup.GET("user-fio/:fio{s}", userRequest.FindByFIO, userHandler.FindByFIO)
		userGroup.GET("user-name/:username{s}", userRequest.FindByUsername, userHandler.FindByUsername)
		userGroup.GET("user/company/:id{n}", userRequest.FindByCompanyId, userHandler.FindByCompanyId)
		userGroup.GET("user/branch/:id{n}", userRequest.FindByBranchId, userHandler.FindByBranchId)
		userGroup.DELETE("user/:userId/role/:roleId", userRequest.DeleteRole, userHandler.DeleteRole)
	}

}
