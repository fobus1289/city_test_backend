package auth

import (
	"github.com/fobus1289/marshrudka/router"
	"v2/app/http/v2/middleware"
	"v2/app/http/v2/request/auth"
)

func Init(drive *router.Drive) {

	authController := &auth.Auth{}

	drive.Dep(authController)

	userGroup := drive.Group("v2/auth")
	{
		userGroup.POST("sign-in", authController.SignIn)
		userGroup.POST("me", middleware.Auth, authController.Me)
	}

}
