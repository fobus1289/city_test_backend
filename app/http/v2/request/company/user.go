package company

import (
	"github.com/fobus1289/marshrudka/router"
	"gorm.io/gorm"
	"net/http"
	"v2/app/http/v2/middleware"
	"v2/app/model"
	"v2/app/service"
)

type User struct {
	*service.Logger
}

// All
// @Security ApiKeyAuth
// @Tags company
// @Description users
// @Param page path integer false "page"
// @Param limit path integer false "limit"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Response
// @Router /v2/company/{id}/user [get]
func (u *User) All(request *router.Request) (page, limit int, scopes middleware.Scopes) {

	page = int(request.QueryGetInt("page"))
	limit = int(request.QueryGetInt("limit"))

	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("company_id", request.Param("companyId"))
	})

	return
}

// FindById
// @Security ApiKeyAuth
// @Tags company
// @Description find user
// @Param id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/company/{id}/user/{id} [get]
func (u *User) FindById(request *router.Request) interface{} {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("company_id = ?", request.Param("companyId")).
			Where("id = ?", request.Param("userId"))
	})
}

// FindByUsername
// @Security ApiKeyAuth
// @Tags company
// @Description find user
// @Param id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/company/{id}/user-name/{username} [get]
func (u *User) FindByUsername(request *router.Request) interface{} {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("company_id = ?", request.Param("companyId")).
			Where("username = ?", request.Param("username"))
	})
}

// FindByFIO
// @Security ApiKeyAuth
// @Tags company
// @Description find user
// @Param id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/company/{id}/user-fio/{fio} [get]
func (u *User) FindByFIO(request *router.Request) interface{} {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("company_id = ?", request.Param("companyId")).
			Where("fio = ?", request.Param("fio"))
	})
}

// FindByBranchId
// @Security ApiKeyAuth
// @Tags company
// @Description find user
// @Param id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/company/{id}/user/branch/{id} [get]
func (u *User) FindByBranchId(request *router.Request) interface{} {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("company_id = ?", request.Param("companyId")).
			Where("branch_id = ?", request.Param("branchId"))
	})
}

// Create
// @Security ApiKeyAuth
// @Tags company
// @Description create user
// @Param input body model.User true "model.User"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/company/{id}/user/ [post]
func (u *User) Create(user *model.User, jwtUser *model.JwtUser, request *router.Request, companyService *service.Company) interface{} {

	if !request.TryParamGetInt("companyId", user.CompanyId) {
		u.Println(request.Param("companyId"), "company id can be empty")
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
		u.Println(request.Param("companyId"), err.Error())
		return router.Response(http.StatusInternalServerError).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	if !exists {
		u.Println(request.Param("companyId"), " company not found")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company not found",
		})
	}

	return user
}

// Update
// @Security ApiKeyAuth
// @Tags company
// @Description update user
// @Param input body model.User true "model.User"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/company/{id}/user/ [put]
func (u *User) Update(user *model.User, companyService *service.Company, jwtUser *model.JwtUser, request *router.Request) interface{} {

	if !request.TryParamGetInt("companyId", user.CompanyId) || !request.TryParamGetInt("userId", &user.Id) {
		u.Println("company id or branch id empty")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company id or branch id empty",
		})
	}

	exists, err := companyService.Exists(func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("id = ?", user.CompanyId).
			Where("user_id =?", jwtUser.Id)
	})

	if err != nil {
		u.Println(request.Param("companyId"), err.Error())
		return router.Response(http.StatusInternalServerError).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	if !exists {
		u.Println(request.Param("companyId"), " company not found")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company not found",
		})
	}

	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("company_id = ?", user.CompanyId).
			Where("id = ?", user.Id)
	})

}

// Delete
// @Security ApiKeyAuth
// @Tags company
// @Description find user
// @Param id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/company/{id}/user/{id} [delete]
func (u *User) Delete(request *router.Request, jwtUser *model.JwtUser, companyService *service.Company) interface{} {

	exists, err := companyService.Exists(func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("id = ?", request.Param("companyId")).
			Where("user_id =?", jwtUser.Id)
	})

	if err != nil {
		u.Println(request.Param("companyId"), err.Error())
		return router.Response(http.StatusInternalServerError).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	if !exists {
		u.Println(request.Param("companyId"), " company not found")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company not found",
		})
	}

	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("id = ?", request.Param("userId")).
			Where("company_id = ?", request.Param("companyId"))
	})

}

// DeleteRole
// @Security ApiKeyAuth
// @Tags company
// @Description find user
// @Param id path integer true "user id"
// @Param id path integer true "role id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/company/{id}/user/{id}/role/{id} [delete]
func (u *User) DeleteRole(request *router.Request, jwtUser *model.JwtUser, companyService *service.Company) interface{} {

	exists, err := companyService.Exists(func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("id = ?", request.Param("companyId")).
			Where("user_id =?", jwtUser.Id)
	})

	if err != nil {
		u.Println(request.Param("companyId"), err.Error())
		return router.Response(http.StatusInternalServerError).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	if !exists {
		u.Println(request.Param("companyId"), " company not found")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "company not found",
		})
	}

	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("user_id = ?", request.Param("userId")).
			Where("role_id = ?", request.Param("roleId"))
	})

}
