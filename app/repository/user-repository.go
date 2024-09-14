package repository

import (
	constant "code/app/constant"
	"code/app/domain/dao"
	"code/app/domain/dto"
	"code/app/pkg"
	"log"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUser(oage int, offset int, search string, filter map[string][]int) ([]dao.User, error)
	FindUserById(id int) (dao.User, error)
	Save(user *dto.CreateUserDTO) (dao.User, error)
	UpdateUserData(dao.User, *dto.UpdateUserDTO) (dao.User, error)
	DeleteUserbyId(id int) int64
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) FindAllUser(limit int, offset int, search string, filter map[string][]int) ([]dao.User, error) {
	var users []dao.User
	// fmt.Printf("page is %d\n", limit)
	// fmt.Printf("offset is %d\n", offset)
	// fmt.Printf("Filter is %q\n", filter)
	// fmt.Printf("Filter is %s\n", search)
	query := u.db
	if search != "" {

		query = query.Where("Name ILIKE  ?", search+"%").Or("Address ILIKE ?", "%"+search+"%")
	}
	// Check if age is provided
	if len(filter) > 0 && len(filter["age"]) > 0 {
		query = query.Where("age in  ?", filter["age"])
	}
	// We can add more fields in this way
	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

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
	if createUserDTO.Birthday != nil {

		birthday, err := time.Parse("2006-01-02", *createUserDTO.Birthday)
		if err != nil {
			return dao.User{}, err
		}
		user.Birthday = &birthday
	}

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

func (u UserRepositoryImpl) DeleteUserbyId(id int) int64 {

	result := u.db.Delete(&dao.User{}, id)
	// fmt.Println(result.RowsAffected)
	rowsAffected := result.RowsAffected
	return rowsAffected
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&dao.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}
