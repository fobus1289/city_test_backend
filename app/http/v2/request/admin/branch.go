package admin

import (
	"github.com/fobus1289/marshrudka/router"
	"gorm.io/gorm"
	"net/http"
	"v2/app/http/v2/middleware"
	"v2/app/model"
	"v2/app/service"
)

type Branch struct {
	*service.Logger
	CompanyService *service.Company
}

// All
// @Security ApiKeyAuth
// @Tags admin branch
// @Description branches
// @Param page path integer false "page"
// @Param limit path integer false "limit"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Response
// @Router /v2/admin/branch/company/{id} [get]
func (b *Branch) All(request *router.Request) (page, limit int, scopes interface{}) {

	page = int(request.QueryGetInt("page"))
	limit = int(request.QueryGetInt("limit"))

	scopes = append(middleware.Scopes{}, func(tx *gorm.DB) *gorm.DB {
		return tx.Where("company_id", request.Param("companyId"))
	})

	return
}

// Create
// @Security ApiKeyAuth
// @Tags admin branch
// @Description create branch
// @Param input body model.Branch true "model.Branch"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Branch
// @Router /v2/admin/branch/{companyId} [post]
func (b *Branch) Create(branch *model.Branch, request *router.Request) interface{} {

	if !request.TryParamGetInt("companyId", branch.CompanyId) {
		b.Println("company id can be empty")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company id can be empty",
		})
	}

	return branch
}

// Update
// @Security ApiKeyAuth
// @Tags admin branch
// @Description update branch
// @Param input body model.Branch true "model.Branch"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Branch
// @Router /v2/admin/branch/{id}/{companyId} [put]
func (b *Branch) Update(branch *model.Branch, request *router.Request) interface{} {

	if !request.TryParamGetInt("companyId", branch.CompanyId) {
		b.Println("company id can be empty")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company id can be empty",
		})
	}

	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", request.Param("id")).
			Where("company_id = ?", request.Param("companyId"))
	})
}

// FindById
// @Security ApiKeyAuth
// @Tags admin branch
// @Description find branch
// @Param id path integer true "branch id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Branch
// @Router /v2/admin/branch/{id} [get]
func (b *Branch) FindById(request *router.Request) (scopes middleware.Scopes) {
	return append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", request.Param("id"))
	})
}

// FindByName
// @Security ApiKeyAuth
// @Tags admin branch
// @Description find branch
// @Param name path integer true "branch name"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Branch
// @Router /v2/admin/branch/{name}/{companyId} [get]
func (b *Branch) FindByName(request *router.Request) (scopes middleware.Scopes) {
	return append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", request.Param("name")).
			Where("company_id = ?", request.Param("companyId"))
	})
}

// FindByUserId
// @Security ApiKeyAuth
// @Tags admin branch
// @Description find branch
// @Param user id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {array} model.Branch
// @Router /v2/admin/branch/user/{id} [get]
func (b *Branch) FindByUserId(request *router.Request) (scopes middleware.Scopes) {
	return append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", request.Param("id"))
	})
}

// Delete
// @Security ApiKeyAuth
// @Tags admin branch
// @Description find branch
// @Param id path integer true "branch id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Branch
// @Router /v2/admin/branch/{id} [delete]
func (b *Branch) Delete(request *router.Request) (scopes middleware.Scopes) {
	return append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", request.Param("id"))
	})
}
