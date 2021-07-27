package handler

import (
	"github.com/fobus1289/marshrudka/router"
	"net/http"
	"v2/app/http/v2/middleware"
	"v2/app/model"
	"v2/app/service"
)

var emptyUser = map[string]interface{}{}

type User struct {
	*service.Logger
	UserService    *service.User
	CompanyService *service.Company
	BranchService  *service.Branch
}

func (u *User) All(page, limit int, scopes middleware.Scopes) interface{} {

	paging, err := u.UserService.Users(page, limit, scopes...)

	if err != nil {
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(paging)
}

func (u *User) Create(user *model.User, jwtUser *model.JwtUser) interface{} {

	if err := u.UserService.Create(user); err != nil {
		u.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(user)
}

func (u *User) Update(user *model.User, scopes middleware.Scopes, jwtUser *model.JwtUser) interface{} {

	if err := u.UserService.Update(user, scopes...); err != nil {
		u.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(user)
}

func (u *User) Delete(scopes middleware.Scopes, jwtUser *model.JwtUser) interface{} {

	if err := u.UserService.Delete(scopes...); err != nil {
		u.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return true
}

func (u *User) FindById(scopes middleware.Scopes, jwtUser *model.JwtUser) interface{} {

	user, err := u.UserService.FindById(scopes...)

	if err != nil {
		u.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	if user.Id == 0 {
		return router.Response(http.StatusBadRequest).Json(emptyUser)
	}

	return router.Response(http.StatusOK).Json(user)
}

func (u *User) FindByCompanyId(scopes middleware.Scopes) interface{} {
	users, err := u.UserService.FindByCompanyId(scopes...)

	if err != nil {
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(users)
}

func (u *User) FindByBranchId(scopes middleware.Scopes) interface{} {
	users, err := u.UserService.FindByBranchId(scopes...)

	if err != nil {
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(users)
}

func (u *User) FindByFIO(scopes middleware.Scopes) interface{} {

	users, err := u.UserService.FindByFIO(scopes...)

	if err != nil {
		u.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(users)
}

func (u *User) FindByUsername(scopes middleware.Scopes) interface{} {

	user, err := u.UserService.FindByUsername(scopes...)

	if err != nil {
		u.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	if user.Id == 0 {
		return router.Response(http.StatusOK).Json(map[string]interface{}{})
	}

	return router.Response(http.StatusOK).Json(user)
}

func (u *User) DeleteRole(scopes middleware.Scopes) interface{} {

	user, err := u.UserService.DeleteRole(scopes...)

	if err != nil {
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	return user
}
