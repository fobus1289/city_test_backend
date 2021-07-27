package auth

import (
	"github.com/fobus1289/marshrudka/router"
	"gorm.io/gorm"
	"net/http"
	"v2/app/model"
	"v2/app/service"
)

type Auth struct {
	*service.Logger
}

// SignIn
// @Security ApiKeyAuth
// @Tags auth
// @Description user login
// @Param input body model.SigInUser true "user info"
// @Produce  json
// @Failure 400 {object} model.ResponseError
// @Success 200 {object} model.JwtUser
// @Router /v2/auth/sign-in [post]
func (a *Auth) SignIn(sigInUser *model.SigInUser, auth *service.Auth) interface{} {

	var (
		err      error
		response map[string]interface{}
	)

	if response, err = auth.SignIn(sigInUser); err != nil {
		a.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(response)
}

// Me
// @Security ApiKeyAuth
// @Tags auth
// @Description get user info
// @Produce  json
// @Failure 401 {object} model.ResponseError
// @Success 200 {object} model.JwtUser
// @Router /v2/auth/me [post]
func (a *Auth) Me(jwtUser *model.JwtUser, userService *service.User) interface{} {

	user, err := userService.FindById(
		func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Companies").Preload("Branches").
				Where("id=?", jwtUser.Id)
		},
	)

	if err != nil {
		a.Println(err)
		return router.Response(http.StatusBadRequest).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	return router.Response(http.StatusOK).Json(user)
}
