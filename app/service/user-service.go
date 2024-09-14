package service

import (
	"code/app/constant"
	"code/app/domain/dto"
	"code/app/pkg"
	"code/app/repository"
	"net/http"
	"strings"

	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUser(c *gin.Context)
	GetUserById(c *gin.Context)
	AddUserData(c *gin.Context)
	UpdateUserData(c *gin.Context)
	//UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (u UserServiceImpl) GetUserById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound, "")
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) AddUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data user")
	var request dto.CreateUserDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		// log.Printf("invalid request value is %d", constant.InvalidRequest)
		pkg.PanicException(constant.InvalidRequest, "")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)

	// Check if email already exist and send proper response

	data, err := u.userRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError, "")
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetAllUser(c *gin.Context) {

	defer pkg.PanicHandler(c)

	// Need to get page and count

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pkg.PanicException(constant.InvalidRequest, "")

	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < 1 {
		pkg.PanicException(constant.InvalidRequest, "")

	}
	offset := (page - 1) * limit

	//Need to get search key
	//Need to get filter key
	search := c.Query("search")
	filter := make(map[string][]int)
	age_str := strings.Split(c.Query("age"), ",")

	var age_int []int

	for _, age := range age_str {

		age_unit, err := strconv.Atoi(age)
		if err != nil {
			pkg.PanicException(constant.InvalidRequest, "")
		}
		age_int = append(age_int, age_unit)
	}
	filter["age"] = age_int

	users, err := u.userRepository.FindAllUser(limit, offset, search, filter)
	if err != nil {
		log.Error("Error during adding data to DB", err)
		pkg.PanicException(constant.UnknownError, "")
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, users))

	// get data as per page and count by DB

}

func (u UserServiceImpl) UpdateUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		pkg.PanicException(constant.InvalidRequest, "")
	}

	user, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound, "")
	}

	var request dto.UpdateUserDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Error binding json with data", err)
		pkg.PanicException(constant.InvalidRequest, "")
	}

	data, err := u.userRepository.UpdateUserData(user, &request)
	if err != nil {
		log.Error("Error during adding data to DB", err)
		pkg.PanicException(constant.UnknownError, "")
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) DeleteUser(c *gin.Context) {
	defer pkg.PanicHandler(c)
	userID, err := strconv.Atoi(c.Param("userID"))

	if err != nil {
		pkg.PanicException(constant.InvalidRequest, "")
	}

	rows := u.userRepository.DeleteUserbyId(userID)

	if rows > 0 {
		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, ""))

	} else {
		pkg.PanicException(constant.DataNotFound, "")
	}

}
func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
