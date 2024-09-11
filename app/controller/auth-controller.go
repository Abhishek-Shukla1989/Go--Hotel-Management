package controller

import (
	"code/app/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Forget(c *gin.Context)
	//Register(c *gin.Context)
}

type AuthControllerImpl struct {
	svc service.AuthService
}

func (u AuthControllerImpl) Login(c *gin.Context) {

	u.svc.LoginUser(c)
}

// func (u AuthControllerImpl) Register(c *gin.Context) {

// }

func (u AuthControllerImpl) Forget(c *gin.Context) {
	u.svc.ForgetService(c)
}

func AuthControllerInit(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		svc: authService,
	}
}
