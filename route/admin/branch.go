package admin

import (
	"github.com/fobus1289/marshrudka/router"
	"v2/app/http/v2/handler"
	"v2/app/http/v2/middleware"
	"v2/app/http/v2/request/admin"
)

func Branch(drive *router.Drive) {

	branchController := &admin.Branch{}
	branchHandler := &handler.Branch{}
	drive.Dep(branchController)
	drive.Dep(branchHandler)

	branchGroup := drive.Group("v2/admin", middleware.Admin)
	{
		branchGroup.GET("branch/company/:id{n}", branchController.All, branchHandler.All)
		branchGroup.GET("branch/:id{n}", branchController.FindById, branchHandler.FindById)
		branchGroup.GET("branch/:name{s}/:companyId{n}", branchController.FindByName, branchHandler.FindByName)
		branchGroup.GET("branch/user/:id{n}", branchController.FindByUserId, branchHandler.FindByUserId)
		branchGroup.DELETE("branch/:id{n}", branchController.Delete, branchHandler.Delete)
		branchGroup.POST("branch/:id{n}", branchController.Create, branchHandler.Create)
		branchGroup.PUT("branch/:id{n}", branchController.Update, branchHandler.Update)
	}

}
