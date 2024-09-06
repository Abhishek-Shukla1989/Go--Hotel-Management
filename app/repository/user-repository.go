package repository

import (
	"code/app/domain/dao"
	"code/app/domain/dto"

	"gorm.io/gorm"
)

type UserRepository interface {
	//FindAllUser() ([]dao.User, error)
	//FindUserById(id int) (dao.User, error)
	Save(user *dto.CreateUserDTO) (dao.User, error)
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

// func (u UserRepositoryImpl) FindUserById(id int) (dao.User, error) {
// 	user := dao.User{
// 		ID: id,
// 	}
// 	err := u.db.Preload("Role").First(&user).Error
// 	if err != nil {
// 		log.Error("Got and error when find user by id. Error: ", err)
// 		return dao.User{}, err
// 	}
// 	return user, nil
// }

func (u UserRepositoryImpl) Save(createUserDTO *dto.CreateUserDTO) (dao.User, error) {

	var user dao.User

	user.Name = createUserDTO.Name
	user.Email = createUserDTO.Email
	user.Password = createUserDTO.Password // You may want to hash the password before saving
	user.Age = createUserDTO.Age
	user.Address = createUserDTO.Address

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
