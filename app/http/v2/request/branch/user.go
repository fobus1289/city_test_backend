package branch

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

// Branch
// @Security ApiKeyAuth
// @Tags branch
// @Description create user
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Branch
// @Router /v2/branch/{id}/ [get]
func (u *User) Branch(request *router.Request, jwtUser *model.JwtUser) interface{} {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("id = ?", request.Param("branchId")).
			Where("user_id = ?", jwtUser.Id)
	})
}

// All
// @Security ApiKeyAuth
// @Tags branch
// @Description users
// @Param page path integer false "page"
// @Param limit path integer false "limit"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.Response
// @Router /v2/branch/{id}/user [get]
func (u *User) All(request *router.Request) (page, limit int, scopes middleware.Scopes) {

	page = int(request.QueryGetInt("page"))
	limit = int(request.QueryGetInt("limit"))

	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("branch_id", request.Param("branchId"))
	})

	return
}

// FindById
// @Security ApiKeyAuth
// @Tags branch
// @Description find user
// @Param id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/branch/{id}/user/{id} [get]
func (u *User) FindById(request *router.Request) interface{} {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("id = ?", request.Param("userId")).
			Where("branch_id = ?", request.Param("branchId"))
	})
}

// FindByUsername
// @Security ApiKeyAuth
// @Tags branch
// @Description find user
// @Param id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/branch/{id}/user-name/{username} [get]
func (u *User) FindByUsername(request *router.Request) interface{} {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("username = ?", request.Param("username")).
			Where("branch_id = ?", request.Param("branchId"))
	})
}

// FindByFIO
// @Security ApiKeyAuth
// @Tags branch
// @Description find user
// @Param id path integer true "user id"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {array} model.User
// @Router /v2/branch/{id}/user-fio/{fio} [get]
func (u *User) FindByFIO(request *router.Request) interface{} {
	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("fio = ?", request.Param("fio")).
			Where("branch_id = ?", request.Param("branchId"))
	})
}

// Create
// @Security ApiKeyAuth
// @Tags branch
// @Description create user
// @Param input body model.User true "model.User"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Failure 401 {object} model.ResponseError
// @Failure 403 {object} model.ResponseError
// @Success 200 {object} model.User
// @Router /v2/branch/{id}/user/ [post]
func (u *User) Create(jwtUser *model.JwtUser, user *model.User, userBranch *service.Branch, request *router.Request) interface{} {

	if !request.TryParamGetInt("branchId", user.BranchId) {
		u.Println("branch id can be empty")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "branch id can be empty",
		})
	}

	has, err := userBranch.Exists(func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("id = ?", user.BranchId).
			Where("user_id = ?", jwtUser.Id)
	})

	if err != nil {
		u.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	if !has {
		u.Println("branch not found: ", user.BranchId)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "branch not found",
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
// @Router /v2/branch/{id}/user/{id} [put]
func (u *User) Update(jwtUser *model.JwtUser, user *model.User, userBranch *service.Branch, request *router.Request) interface{} {

	if !request.TryParamGetInt("branchId", user.BranchId) {
		u.Println("branch id can be empty")
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "branch id can be empty",
		})
	}

	has, err := userBranch.Exists(func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("id = ?", user.BranchId).
			Where("user_id = ?", jwtUser.Id)
	})

	if err != nil {
		u.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	if !has {
		u.Println("branch not found: ", user.BranchId)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": "branch not found",
		})
	}

	return append(middleware.Scopes{}, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("id = ?", user.BranchId).
			Where("user_id = ?", jwtUser.Id)
	})

}
