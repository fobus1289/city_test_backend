package handler

import (
	"github.com/fobus1289/marshrudka/router"
	"net/http"
	"v2/app/http/v2/middleware"
	"v2/app/model"
	"v2/app/service"
)

type Company struct {
	*service.Logger
	UserService    *service.User
	CompanyService *service.Company
	BranchService  *service.Branch
}

func NewCompany() *Company {
	return &Company{}
}

func (c *Company) All(page, limit int, scopes middleware.Scopes) interface{} {

	paging, err := c.CompanyService.Companies(page, limit, scopes...)

	if err != nil {
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(paging)
}

func (c *Company) Create(company *model.Company) interface{} {

	company, err := c.CompanyService.Create(company)

	if err != nil {
		c.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(company)
}

func (c *Company) Update(company *model.Company, scopes middleware.Scopes) interface{} {

	company, err := c.CompanyService.Update(company, scopes...)

	if err != nil {
		c.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(company)
}

func (c *Company) Delete(scopes middleware.Scopes, jwtUser *model.JwtUser) interface{} {

	if err := c.CompanyService.Delete(scopes...); err != nil {
		c.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return true
}

func (c *Company) FindById(scopes middleware.Scopes) interface{} {

	company, err := c.CompanyService.FindById(scopes...)

	if err != nil {
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(company)
}

func (c *Company) FindByName(scopes middleware.Scopes) interface{} {

	companies, err := c.CompanyService.FindByName(scopes...)

	if err != nil {
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(companies)
}

func (c *Company) FindByINN(scopes middleware.Scopes) interface{} {

	company, err := c.CompanyService.FindByINN(scopes...)

	if err != nil {
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(company)
}

func (c *Company) FindByUserId(scopes middleware.Scopes) interface{} {

	companies, err := c.CompanyService.FindByUserId(scopes...)

	if err != nil {
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(companies)
}

func (c *Company) FindByBranchId(scopes middleware.Scopes) interface{} {

	company, err := c.CompanyService.FindByBranchId(scopes...)

	if err != nil {
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(company)
}
