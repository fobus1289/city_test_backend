package admin

import (
	"github.com/fobus1289/marshrudka/router"
	"gorm.io/gorm"
	"v2/app/http/v2/middleware"
	"v2/app/model"
	"v2/app/service"
)

type Company struct {
	*service.Logger
}

// Companies
// @Security ApiKeyAuth
// @Tags admin company
// @Description users
// @Param page path integer false "page"
// @Param limit path integer false "limit"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Response
// @Router /v2/admin/companies/ [get]
func (c *Company) Companies(request *router.Request) (page, limit int, scopes middleware.Scopes) {
	page = int(request.ParamGetInt("page"))
	limit = int(request.ParamGetInt("limit"))
	return
}

// FindById
// @Security ApiKeyAuth
// @Tags admin company
// @Description find company
// @Param id path integer true "company id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Company
// @Router /v2/company/{id} [get]
func (c *Company) FindById(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", request.Param("id"))
	})
}

// Create
// @Security ApiKeyAuth
// @Tags admin company
// @Description find company
// @Param input body model.Company true "model.Company"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Company
// @Router /v2/company [post]
func (c *Company) Create(company *model.Company) *model.Company {
	return company
}

// Update
// @Security ApiKeyAuth
// @Tags admin company
// @Description find company
// @Param input body model.Company true "model.Company"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Company
// @Router /v2/company/{id} [put]
func (c *Company) Update(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", request.Param("id"))
	})
}

// Delete
// @Security ApiKeyAuth
// @Tags admin company
// @Description find company
// @Param id path integer true "company id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Company
// @Router /v2/company/{id} [delete]
func (c *Company) Delete(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", request.Param("id"))
	})
}

// FindByName
// @Security ApiKeyAuth
// @Tags admin company
// @Description find company
// @Param name path integer true "company name"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Company
// @Router /v2/company/{name} [get]
func (c *Company) FindByName(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+request.Param("name")+"%")
	})
}

// FindByUserId
// @Security ApiKeyAuth
// @Tags admin company
// @Description find company
// @Param id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Company
// @Router /v2/company/user/{id} [get]
func (c *Company) FindByUserId(request *router.Request) middleware.Scopes {

	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", request.Param("id"))
	})

}

// FindByBranchId
// @Security ApiKeyAuth
// @Tags admin company
// @Description find company
// @Param id path integer true "branch id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Company
// @Router /v2/company/branch/{id} [get]
func (c *Company) FindByBranchId(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", request.Param("id"))
	})
}
