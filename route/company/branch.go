package company

import (
	"github.com/fobus1289/marshrudka/router"
	"v2/app/http/v2/handler"
	"v2/app/http/v2/middleware"
	"v2/app/http/v2/request/company"
)

func Branch(drive *router.Drive) {

	branchRequest := &company.Branch{}
	branchHandler := &handler.Branch{}

	drive.Dep(branchRequest)
	drive.Dep(branchHandler)

	branchGroup := drive.Group("v2/company/:companyId{n}", middleware.Company)
	{
		branchGroup.GET("branch", branchRequest.All, branchHandler.All)
		branchGroup.GET("branch/:name{s}", branchRequest.FindByName, branchHandler.FindByName)
		branchGroup.GET("branch/user/:userId{n}", branchRequest.FindByUserId, branchHandler.FindByUserId)
		branchGroup.GET("branch/:branchId{n}", branchRequest.FindByName, branchHandler.FindByName)
		branchGroup.POST("branch", branchRequest.Create, branchHandler.Create)
		branchGroup.PUT("branch/:branchId{n}", branchRequest.Update, branchHandler.Update)
		branchGroup.DELETE("branch/:branchId{n}", branchRequest.FindByName, branchHandler.FindByName)
	}

}
