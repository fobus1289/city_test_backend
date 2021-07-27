package handler

import (
	"github.com/fobus1289/marshrudka/router"
	"net/http"
	"v2/app/http/v2/middleware"
	"v2/app/model"
	"v2/app/service"
)

type Branch struct {
	*service.Logger
	UserService    *service.User
	CompanyService *service.Company
	BranchService  *service.Branch
}

func (b *Branch) All(page, limit int, scopes middleware.Scopes) interface{} {

	paging, err := b.BranchService.Branches(page, limit, scopes...)

	if err != nil {
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(paging)
}

func (b *Branch) Create(branch *model.Branch) interface{} {

	branch, err := b.BranchService.Create(branch)

	if err != nil {
		b.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(branch)
}

func (b *Branch) Update(branch *model.Branch, scopes middleware.Scopes) interface{} {

	branch, err := b.BranchService.Update(branch, scopes...)

	if err != nil {
		b.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(branch)
}

func (b *Branch) Delete(scopes middleware.Scopes, jwtUser *model.JwtUser) interface{} {

	if err := b.BranchService.Delete(scopes...); err != nil {
		b.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return true
}

func (b *Branch) FindById(scopes middleware.Scopes, jwtUser *model.JwtUser) interface{} {

	branch, err := b.BranchService.FindById(scopes...)

	if err != nil {
		b.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(branch)
}

func (b *Branch) FindByUserId(scopes middleware.Scopes, jwtUser *model.JwtUser) interface{} {

	branches, err := b.BranchService.FindByUserId(scopes...)

	if err != nil {
		b.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(branches)
}

func (b *Branch) FindByCompanyId(scopes middleware.Scopes, jwtUser *model.JwtUser) interface{} {

	branches, err := b.BranchService.FindByCompanyId(scopes...)

	if err != nil {
		b.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(branches)
}

func (b *Branch) FindByName(scopes middleware.Scopes, jwtUser *model.JwtUser) interface{} {

	branches, err := b.BranchService.FindByName(scopes...)

	if err != nil {
		b.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(branches)
}
