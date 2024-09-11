package service

import (
	"code/app/constant"
	"code/app/domain/dto"
	"code/app/pkg"
	"code/app/repository"
	"code/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	LoginUser(c *gin.Context)
	ForgetService(c *gin.Context)
}

type AuthServiceImpl struct {
	authRepository repository.AuthRepository
}

func (a AuthServiceImpl) LoginUser(c *gin.Context) {

	defer pkg.PanicHandler(c)

	var request dto.LoginUserDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest, "")
	}
	// check user exist or not or password is wrong
	user, err := a.authRepository.Login(&request)
	if err != nil {
		pkg.PanicException(constant.InvalidRequest, "")

	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		pkg.PanicException(constant.UnknownError, "")
	}

	data := map[string]interface{}{

		"user":  user,
		"token": token,
	}

	// generate JWT token with email
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (a AuthServiceImpl) ForgetService(c *gin.Context) {

	defer pkg.PanicHandler(c)

	var request dto.ForgetPassword
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest, "")
	}
	// check

	data, err := a.authRepository.ForgetPassword(&request)
	if err != nil {
		pkg.PanicException(constant.UnknownError, "")
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))

}
func AuthServiceInit(authRepository repository.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepository: authRepository,
	}
}
