package admin

import (
	"github.com/fobus1289/marshrudka/router"
	"v2/app/http/v2/handler"
	"v2/app/http/v2/middleware"
	"v2/app/http/v2/request/admin"
)

func Company(drive *router.Drive) {

	companyRequest := &admin.Company{}
	companyHandler := &handler.Company{}

	drive.Dep(companyRequest)
	drive.Dep(companyHandler)

	companyGroup := drive.Group("v2/admin", middleware.Admin)
	{
		companyGroup.GET("companies", companyRequest.Companies, companyHandler.All)
		companyGroup.GET("company/:id{n}", companyRequest.FindById, companyHandler.FindById)
		companyGroup.POST("company", companyRequest.Create, companyHandler.Create)
		companyGroup.PUT("company", companyRequest.Update, companyHandler.Update)
		companyGroup.DELETE("company/:id{n}", companyRequest.Delete, companyHandler.Delete)
		companyGroup.DELETE("company/:name{s}", companyRequest.FindByName, companyHandler.FindByName)
		companyGroup.DELETE("company/user/:id{n}", companyRequest.FindByUserId, companyHandler.FindByUserId)
		companyGroup.DELETE("company/branch/:id{n}", companyRequest.FindByBranchId, companyHandler.FindByBranchId)
	}

}
