package company

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
}

// All
// @Security ApiKeyAuth
// @Tags company
// @Description branches
// @Param page path integer false "page"
// @Param limit path integer false "limit"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Response
// @Router /v2/company/{id}/branch [get]
func (b *Branch) All(request *router.Request) (page, limit int, scopes middleware.Scopes) {

	page = int(request.QueryGetInt("page"))
	limit = int(request.QueryGetInt("limit"))

	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("company_id", request.Param("companyId"))
	})

	return
}

// FindByName
// @Security ApiKeyAuth
// @Tags company
// @Description find branch
// @Param name path integer true "branch name"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {array} model.Branch
// @Router /v2/company/{id}/branch/{name} [get]
func (b *Branch) FindByName(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("company_id = ?", request.Param("companyId")).
			Where("name like ?", "%"+request.Param("name")+"%")
	})
}

// FindByUserId
// @Security ApiKeyAuth
// @Tags company
// @Description find branch
// @Param user id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {array} model.Branch
// @Router /v2/company/{id}/branch/user/{id} [get]
func (b *Branch) FindByUserId(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("company_id = ?", request.Param("companyId")).
			Where("user_id = ?", request.Param("userId"))
	})
}

// FindById
// @Security ApiKeyAuth
// @Tags company
// @Description find branch
// @Param id path integer true "branch id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Branch
// @Router /v2/company/{id}/branch/{id} [get]
func (b *Branch) FindById(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("company_id = ?", request.Param("companyId")).
			Where("id = ?", request.Param("branchId"))
	})
}

// Create
// @Security ApiKeyAuth
// @Tags company
// @Description create branch
// @Param input body model.Branch true "model.Branch"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Success 200 {object} model.Branch
// @Router /v2/company/{id}/branch [post]
func (b *Branch) Create(request *router.Request, branch *model.Branch, jwtUser *model.JwtUser, companyService *service.Company) interface{} {

	if !request.TryParamGetInt("companyId", branch.CompanyId) {
		b.Println(request.Param("companyId"), "company id can be empty")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company id can be empty",
		})
	}

	exists, err := companyService.Exists(func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("id =?", request.Param("companyId")).
			Where("user_id = ?", jwtUser.Id)
	})

	if err != nil {
		b.Println(request.Param("companyId"), err.Error())
		return router.Response(http.StatusInternalServerError).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	if !exists {
		b.Println(request.Param("companyId"), " company not found")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company not found",
		})
	}

	return branch
}

// Update
// @Security ApiKeyAuth
// @Tags company
// @Description update branch
// @Param input body model.Branch true "model.Branch"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Success 200 {object} model.Branch
// @Router /v2/company/{id}/branch/{id} [put]
func (b *Branch) Update(branch *model.Branch, companyService *service.Company, jwtUser *model.JwtUser, request *router.Request) interface{} {

	if !request.TryParamGetInt("companyId", branch.CompanyId) || !request.TryParamGetInt("branchId", &branch.Id) {
		b.Println("company id or branch id empty")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company id or branch id empty",
		})
	}

	exists, err := companyService.Exists(func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("id = ?", branch.CompanyId).
			Where("user_id =?", jwtUser.Id)
	})

	if err != nil {
		b.Println(request.Param("companyId"), err.Error())
		return router.Response(http.StatusInternalServerError).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	if !exists {
		b.Println(request.Param("companyId"), " company not found")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company not found",
		})
	}

	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("company_id = ?", branch.CompanyId).
			Where("id = ?", branch.Id)
	})
}

// Delete
// @Security ApiKeyAuth
// @Tags company
// @Description update branch
// @Param input body model.Branch true "model.Branch"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Success 200 {object} model.Branch
// @Router /v2/company/{id}/branch/{id} [delete]
func (b *Branch) Delete(request *router.Request, jwtUser *model.JwtUser, companyService *service.Company) interface{} {

	exists, err := companyService.Exists(func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("id = ?", request.Param("companyId")).
			Where("user_id =?", jwtUser.Id)
	})

	if err != nil {
		b.Println(request.Param("branchId"), err.Error())
		return router.Response(http.StatusInternalServerError).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	if !exists {
		b.Println(request.Param("branchId"), " company not found")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company not found",
		})
	}

	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("id = ?", request.Param("branchId")).
			Where("company_id = ?", request.Param("companyId"))
	})
}
