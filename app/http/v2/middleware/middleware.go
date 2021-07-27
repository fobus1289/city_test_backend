package middleware

import (
	"errors"
	"github.com/fobus1289/marshrudka/router"
	"gorm.io/gorm"
	"net/http"
	"v2/app/model"
	"v2/app/service"
)

const (
	forbidden     = "Forbidden denied"
	unauthorized  = "Unauthorized"
	authorization = "Authorization"
)

var (
	unauthorizedError = errors.New(unauthorized)
	forbiddenError    = errors.New(forbidden)
)

type Scopes []func(db *gorm.DB) *gorm.DB

func Cross(w http.ResponseWriter, r *http.Request) interface{} {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Accept", "*/*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Accept-Encoding", "gzip, deflate, br")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return router.Throw{
			StatusCode: 204,
		}
	}

	return nil
}

func Auth(r *http.Request, auth *service.Auth, logger *service.Logger) interface{} {

	user, err := dry(r, auth)

	if err != nil {
		logger.Println(err)
		return router.Response(http.StatusUnauthorized).Throw().Json(map[string]string{
			"message": err.Error(),
		})
	}

	return user
}

func Admin(r *http.Request, auth *service.Auth, logger *service.Logger) interface{} {

	user, err := dry(r, auth)

	if err != nil {
		logger.Println(err)
		return router.Response(http.StatusForbidden).Throw().Json(map[string]string{
			"message": forbiddenError.Error(),
		})
	}

	if !user.Roles.HasRole("admin") {
		logger.Println(forbidden)
		return router.Response(http.StatusForbidden).Throw().Json(map[string]string{
			"message": forbidden,
		})
	}

	return user
}

func Company(r *http.Request, auth *service.Auth, logger *service.Logger) interface{} {

	user, err := dry(r, auth)

	if err != nil {
		logger.Println(err)

		return router.Response(http.StatusForbidden).Throw().Json(map[string]string{
			"message": forbiddenError.Error(),
		})
	}

	if !user.Roles.HasRole("director") {
		logger.Println(forbidden)

		return router.Response(http.StatusForbidden).Throw().Json(map[string]string{
			"message": forbidden,
		})
	}

	return user
}

func Branch(r *http.Request, auth *service.Auth, logger *service.Logger) interface{} {

	user, err := dry(r, auth)

	if err != nil {
		logger.Println(err)
		return router.Response(http.StatusForbidden).Throw().Json(map[string]string{
			"message": forbiddenError.Error(),
		})
	}

	if !user.Roles.HasRole("manager") {
		logger.Println(err)
		return router.Response(http.StatusForbidden).Throw().Json(map[string]string{
			"message": forbidden,
		})
	}

	return user
}

func dry(r *http.Request, auth *service.Auth) (*model.JwtUser, error) {

	header := r.Header.Get(authorization)

	if header == "" || len(header) < 8 {
		return nil, unauthorizedError
	}

	user, err := auth.Decode(header[7:])

	if err != nil {
		return nil, unauthorizedError
	}

	return user, nil
}
