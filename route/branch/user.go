package branch

import (
	"github.com/fobus1289/marshrudka/router"
	"v2/app/http/v2/handler"
	"v2/app/http/v2/middleware"
	"v2/app/http/v2/request/branch"
)

func User(drive *router.Drive) {

	userRequest := &branch.User{}
	userHandler := &handler.User{}
	drive.Dep(userRequest)
	drive.Dep(userHandler)

	userGroup := drive.Group("v2/branch/:branchId{n}", middleware.Company)
	{
		userGroup.GET("/", userRequest.All)
		userGroup.GET("user", userRequest.All, userHandler.All)
		userGroup.GET("user/:userId{n}", userRequest.FindById, userHandler.FindById)
		userGroup.GET("user-name/:username{s}", userRequest.FindByUsername, userHandler.FindByUsername)
		userGroup.GET("user-fio/:fio{s}", userRequest.FindByFIO, userHandler.FindByFIO)
		userGroup.POST("user", userRequest.Create, userHandler.Create)
		userGroup.PUT("user/:userId{n}", userRequest.Update, userHandler.Update)
	}

}
