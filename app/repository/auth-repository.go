package repository

import (
	"code/app/constant"
	"code/app/domain/dao"
	"code/app/domain/dto"
	"code/app/pkg"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(loginDTO *dto.LoginUserDTO) (dao.User, error)
	ForgetPassword(forgetPassDTO *dto.ForgetPassword) (bool, error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func (a AuthRepositoryImpl) Login(loginDTO *dto.LoginUserDTO) (dao.User, error) {

	var existigUser dao.User

	if err := a.db.Where("email = ?", loginDTO.Email).First(&existigUser).Error; err != nil {
		pkg.PanicException(constant.DataNotFound, "")

	}
	// match password first

	if err := bcrypt.CompareHashAndPassword([]byte(existigUser.Password), []byte(loginDTO.Password)); err != nil {
		pkg.PanicException(constant.DataNotFound, "")
	}
	return existigUser, nil

}

func (a AuthRepositoryImpl) ForgetPassword(loginDTO *dto.ForgetPassword) (bool, error) {

	var existigUser dao.User

	if err := a.db.Where("email = ?", loginDTO.Email).First(&existigUser).Error; err != nil {
		pkg.PanicException(constant.DataNotFound, "")

	}
	// match concept of sending email to this account and return error

	return true, nil

}

func AuthRepositoryInit(db *gorm.DB) *AuthRepositoryImpl {
	db.AutoMigrate(&dao.User{})
	return &AuthRepositoryImpl{
		db: db,
	}
}
