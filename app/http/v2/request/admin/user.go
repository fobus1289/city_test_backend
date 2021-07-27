package admin

import (
	"github.com/fobus1289/marshrudka/router"
	"gorm.io/gorm"
	"v2/app/http/v2/middleware"
	"v2/app/model"
	"v2/app/service"
)

type User struct {
	*service.Logger
}

// All
// @Security ApiKeyAuth
// @Tags admin user
// @Description users
// @Param page path integer false "page"
// @Param limit path integer false "limit"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Response
// @Router /v2/admin/users/ [get]
func (u *User) All(request *router.Request) (page, limit int, scopes middleware.Scopes) {
	page = int(request.QueryGetInt("page"))
	limit = int(request.QueryGetInt("limit"))
	return
}

// FindById
// @Security ApiKeyAuth
// @Tags admin user
// @Description find user
// @Param id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/admin/user/{id} [get]
func (u *User) FindById(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", request.Param("id"))
	})
}

// FindByFIO
// @Security ApiKeyAuth
// @Tags admin user
// @Description find user
// @Param fio path integer true "user fio"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {array} model.User
// @Router /v2/admin/user-fio/{fio} [get]
func (u *User) FindByFIO(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("fio like ?", "%"+request.Param("fio")+"%")
	})
}

// FindByUsername
// @Security ApiKeyAuth
// @Tags admin user
// @Description find user
// @Param fio path integer true "user fio"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/admin/user-name/{name} [get]
func (u *User) FindByUsername(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("username = ?", request.Param("username"))
	})
}

// Create
// @Security ApiKeyAuth
// @Tags admin user
// @Description create user
// @Param input body model.User true "model.User"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/admin/user/ [post]
func (u *User) Create(user *model.User) *model.User {
	return user
}

// Update
// @Security ApiKeyAuth
// @Tags admin user
// @Description update user
// @Param input body model.User true "model.User"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/admin/user/{id} [put]
func (u *User) Update(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", request.Param("id"))
	})
}

// Delete
// @Security ApiKeyAuth
// @Tags admin user
// @Description delete user
// @ID Delete
// @Param id path integer true "user id"
// @Produce  plain
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {integer} int
// @Router /v2/admin/user/{id} [delete]
func (u *User) Delete(request *router.Request) middleware.Scopes {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", request.Param("id"))
	})
}

// FindByCompanyId
// @Security ApiKeyAuth
// @Tags admin user
// @Description company users
// @ID FindByCompanyUsers
// @Param company id path integer true "company id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {array} model.User
// @Router /v2/admin/user/company/{id} [get]
func (u *User) FindByCompanyId(request *router.Request) interface{} {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("company_id = ?", request.Param("id"))
	})
}

// FindByBranchId
// @Security ApiKeyAuth
// @Tags admin user
// @Description branch users
// @ID FindByBranchUsers
// @Param branch id path integer true "branch id"
// @Produce json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {array} model.User
// @Router /v2/admin/user/branch/{id} [get]
func (u *User) FindByBranchId(request *router.Request) interface{} {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("branch_id = ?", request.Param("id"))
	})
}

// DeleteRole
// @Security ApiKeyAuth
// @Tags admin user
// @Description delete user role
// @ID DeleteRole
// @Param user id path integer true "user id"
// @Param role id path integer true "role id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {array} model.User
// @Router /v2/admin/user/{id}/role/{id} [delete]
func (u *User) DeleteRole(request *router.Request) interface{} {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", request.Param("userId")).
			Where("role_id = ?", request.Param("roleId"))
	})
}
