package repository

import (
	constant "code/app/constant"
	"code/app/domain/dao"
	"code/app/domain/dto"
	"code/app/pkg"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	//FindAllUser() ([]dao.User, error)
	FindUserById(id int) (dao.User, error)
	Save(user *dto.CreateUserDTO) (dao.User, error)
	UpdateUserData(dao.User, *dto.UpdateUserDTO) (dao.User, error)

	//DeleteUserById(id int) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// func (u UserRepositoryImpl) FindAllUser() ([]dao.User, error) {
// 	var users []dao.User

// 	var err = u.db.Preload("Role").Find(&users).Error
// 	if err != nil {
// 		log.Error("Got an error finding all couples. Error: ", err)
// 		return nil, err
// 	}

// 	return users, nil
// }

func (u UserRepositoryImpl) FindUserById(id int) (dao.User, error) {

	user := dao.User{
		ID: uint(id),
	}
	log.Default().Printf("User data is %v\n\n\n", user)

	if err := u.db.Select("name", "age", "email", "birthday", "address").First(&user).Error; err != nil {
		return dao.User{}, err
	}
	log.Default().Printf("User data is %v", user)
	return user, nil

}
func (u UserRepositoryImpl) Save(createUserDTO *dto.CreateUserDTO) (dao.User, error) {

	var existingUser dao.User

	if err := u.db.Where("email = ?", createUserDTO.Email).First(&existingUser).Error; err == nil {
		pkg.PanicException(constant.ResourceAlreadyExists, "Email")
	}

	var user dao.User

	user.Name = createUserDTO.Name
	user.Email = createUserDTO.Email
	user.Password = createUserDTO.Password // You may want to hash the password before saving
	user.Age = createUserDTO.Age
	user.Address = createUserDTO.Address
	//user.Birthday = createUserDTO.Birthday

	// // Convert the string Birthday to time.Time (if provided)
	// if createUserDTO.Birthday != nil {
	// 	birthday, err := time.Parse("2006-01-02", *createUserDTO.Birthday)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	user.Birthday = &birthday
	// }

	// Return the saved user
	// Save the user using GORM's Create method

	if err := u.db.Create(&user).Error; err != nil {
		return dao.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) UpdateUserData(user dao.User, userData *dto.UpdateUserDTO) (dao.User, error) {

	user.Name = userData.Name
	user.Age = userData.Age
	user.Address = userData.Address
	//user.Birthday = userData.Birthday

	if err := u.db.Save(&user).Error; err != nil {
		return dao.User{}, err
	}
	return user, nil
}

// func (u UserRepositoryImpl) DeleteUserById(id int) error {
// 	err := u.db.Delete(&dao.User{}, id).Error
// 	if err != nil {
// 		log.Error("Got an error when delete user. Error: ", err)
// 		return err
// 	}
// 	return nil
// }

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&dao.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}
