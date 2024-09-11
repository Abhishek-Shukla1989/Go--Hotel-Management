package config

import (
	"code/app/controller"
	"code/app/repository"
	"code/app/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController
	authRepo repository.AuthRepository
	authSvc  service.AuthService
	AuthCtrl controller.AuthController
	// RoleRepo repository.RoleRepository
}

func NewInitialization(userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	authRepo repository.AuthRepository,
	authSvc service.AuthService,
	AuthCtrl controller.AuthController,
) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
		authRepo: authRepo,
		authSvc:  authSvc,
		AuthCtrl: AuthCtrl,
		// RoleRepo: roleRepo,
	}
}
